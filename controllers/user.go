package controllers

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/lib/pq"
	"nino.sh/api/graphql/types"
	"nino.sh/api/utils"
)

func (c *Controller) GetUsers(
	ctx context.Context,
	connection *sql.DB,
) ([]*types.User, error) {
	stmt, err := connection.PrepareContext(ctx, `
		select user_id, language, prefixes from users
    `); if err != nil {
    	return nil, err
	}

	var users []*types.User
    rows, err := stmt.QueryContext(ctx); if err != nil {
    	return nil, err
	}

	defer utils.SwallowError(rows)
    for rows.Next() {
		var (
			userID string
			language string
			prefixes []string
		)

		err = rows.Scan(&userID, &language, pq.Array(&prefixes)); if err != nil {
			return nil, err
		}

		users = append(users, &types.User{
			Prefixes: prefixes,
			Language: language,
			ID: userID,
		})
	}

	return users, nil
}

func (c *Controller) GetUser(
	ctx context.Context,
	conn *sql.DB,
	id string,
) (*types.User, error) {
	users, err := c.GetUsers(ctx, conn); if err != nil {
		return nil, err
	}

	var user *types.User
	for _, u := range users {
		if u.ID == id {
			user = u
			break
		}
	}

	return user, nil
}

func (c *Controller) AddUserPrefix(
	context context.Context,
	conn *sql.DB,
	id string,
	prefix string,
) (bool, error) {
	user, err := c.GetUser(context, conn, id); if err != nil {
		return false, err
	}

	if userPrefixExists(user, prefix) {
		return false, errors.New(fmt.Sprintf("prefix %s already exists in the db", prefix))
	}

	stmt, err := conn.PrepareContext(context, `
		UPDATE guilds SET prefixes = array_append(prefixes, $1) WHERE guild_id = $2
	`); if err != nil {
		return false, err
	}

	_, err = stmt.Query(prefix, user.ID)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (c *Controller) RemoveUserPrefix(
	context context.Context,
	conn *sql.DB,
	id string,
	prefix string,
) (bool, error) {
	user, err := c.GetUser(context, conn, id); if err != nil {
		return false, err
	}

	if !userPrefixExists(user, prefix) {
		return false, errors.New(fmt.Sprintf("prefix %s was not a prefix", prefix))
	}

	stmt, err := conn.PrepareContext(context, `
		UPDATE users SET prefixes = array_remove(prefixes, $1) WHERE guild_id = $2
	`); if err != nil {
		return false, err
	}

	_, err = stmt.Query(prefix, user.ID)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (c *Controller) UpdateUserLanguage(
	context context.Context,
	conn *sql.DB,
	id string,
	language string,
) (bool, error) {
	user, err := c.GetUser(context, conn, id); if err != nil {
		return false, err
	}

	for _, lang := range utils.Languages() {
		if language != lang {
			return false, errors.New(fmt.Sprintf("language %s was not found.", language))
		}
	}

	stmt, err := conn.PrepareContext(context, `
		UPDATE users SET language = $1 WHERE user_id = $2
	`); if err != nil {
		return false, err
	}

	_, err = stmt.Query(language, user.ID)
	if err != nil {
		return false, err
	}

	return true, nil
}

func userPrefixExists(user *types.User, prefix string) bool {
	for _, pre := range user.Prefixes {
		if prefix == pre {
			return true
		}
	}

	return false
}
