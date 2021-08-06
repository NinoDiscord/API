package resolvers

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
	"os"
)

func (r *Resolver) CheckAuthorization(token string) error {
	signingKey := []byte(os.Getenv("SIGNING_KEY"))
	t, err := jwt.Parse(token, func (token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("cannot detect jwt")
		}

		return signingKey, nil
	}); if err != nil {
		return err
	}

	if t.Valid {
		if err := r.Redis.Connection.HGet(context.TODO(), "nino:sessions", t.Raw).Err(); err != nil {
			if err == redis.Nil {
				return errors.New("token is invalid, please re-login! :D")
			}

			return err
		}

		return nil
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors & (jwt.ValidationErrorExpired | jwt.ValidationErrorNotValidYet) != 0 {
			// invalidate token in redis, if possible!
			err := r.Redis.Connection.HDel(context.TODO(), "nino:sessions", t.Raw).Err(); if err != nil {
				// if it was nil, it's possible that it didn't exist in the first place
				if err != redis.Nil {
					logrus.WithField("middleware", "Authentication").Fatalf("unable to invalidate token %s: %v", t.Raw, err)
				}

				return err
			}

			return ve.Inner
		}
	} else {
		return errors.New("possible that jwt is corrupt, re-login. :c")
	}

	return errors.New("unable to detect jwt, create issue: https://github.com/NinoDiscord/API/issues")
}
