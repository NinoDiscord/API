package resolvers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
	"net/http"
	"nino.sh/api/graphql/types"
	"nino.sh/api/utils"
	"os"
	"time"
)

type User struct {
	Discriminator string  `json:"discriminator"`
	PremiumType   int32   `json:"premium_type"`
	PublicFlags   int32   `json:"public_flags"`
	MfaEnabled    *bool   `json:"mfa_enabled"`
	Username      string  `json:"username"`
	Avatar        *string `json:"avatar"`
	System        *bool   `json:"system"`
	Locale        *string `json:"locale"`
	Flags         int32   `json:"flags"`
	ID            string  `json:"id"`
}

type TokenizedUser struct {
	Guilds []types.PartialGuild
	Entry *types.User
	User  User
}

func (u *TokenizedUser) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, u)
}

func (u *TokenizedUser) MarshalBinary() ([]byte, error) {
	return json.Marshal(u)
}

func (r *Resolver) Login(ctx context.Context, args struct{ AccessToken string }) (*string, error) {
	var signingKey = []byte(os.Getenv("SIGNING_KEY"))

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	client := http.Client{}
	req, err := http.NewRequest("GET", "https://discord.com/api/v9/users/@me", nil); if err != nil {
		return nil, err
	}

	req.Header = http.Header{
		"Authorization": []string{fmt.Sprintf("Bearer %s", args.AccessToken)},
	}

	resp, err := client.Do(req); if err != nil {
		return nil, err
	}

	defer utils.SwallowHttpError(resp.Body)
	var user User

	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}

	// Return a list of the user's guilds
	req, err = http.NewRequest("GET", "https://discord.com/api/v9/users/@me/guilds", nil); if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", args.AccessToken))
	resp, err = client.Do(req); if err != nil {
		return nil, err
	}
	defer utils.SwallowHttpError(resp.Body)

	var guilds []types.PartialGuild
	if err := json.NewDecoder(resp.Body).Decode(&guilds); err != nil {
		return nil, err
	}

	// get db entry
	entry, err := r.Controller.GetUser(ctx, r.Db.Connection, user.ID); if err != nil {
		return nil, err
	}

	claims["entry"] = entry
	claims["user"] = user
	claims["exp"] = time.Now().UTC().Add(time.Hour * 48).Unix()

	t, err := token.SignedString(signingKey); if err != nil {
		return nil, err
	}

	// insert it to redis
	if _, err := r.Redis.Connection.HSet(ctx, "nino:sessions", t, &TokenizedUser{
		User: user,
		Entry: entry,
		Guilds: guilds,
	}).Result(); err != nil {
		return nil, err
	}

	// TODO: add expiration instead of manually deleting it on request
	return &t, nil
}

func (r *Resolver) Logout(ctx context.Context) bool {
	token := ctx.Value("token").(string)
	if err := r.CheckAuthorization(token); err != nil {
		return false
	}

	// invalidate it in redis
	if err := r.Redis.Connection.HDel(ctx, "nino:sessions", token).Err(); err != nil {
		logrus.Fatalf("Unable to invalidate token %s - %v", token, err)
		return false
	}

	return true
}

func (r *Resolver) Me(ctx context.Context) (*types.LoggedInUser, error) {
	token := ctx.Value("token").(string)
	if err := r.CheckAuthorization(token); err != nil {
		return nil, err
	}

	// get token value
	value := r.Redis.Connection.HGet(ctx, "nino:sessions", token); if value.Err() != nil {
		if value.Err() == redis.Nil {
			return nil, nil
		}

		err := value.Err()
		return nil, err
	}

	var loggedIn *TokenizedUser
	if err := json.Unmarshal([]byte(value.Val()), &loggedIn); err != nil {
		return nil, err
	}

	return &types.LoggedInUser{
		Discriminator: loggedIn.User.Discriminator,
		Username:      loggedIn.User.Username,
		Avatar:        loggedIn.User.Avatar,
		Guilds:        loggedIn.Guilds,
		Entry:         *loggedIn.Entry,
		ID: 		   loggedIn.User.ID,
	}, nil
}
