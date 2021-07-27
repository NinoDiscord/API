package controllers

import (
	"context"
	"database/sql"
	"github.com/lib/pq"
	"nino.sh/api/graphql/types"
	"nino.sh/api/utils"
)

func (c *Controller) GetUsers(
	ctx context.Context,
	connection *sql.DB,
) ([]*types.User, error) {
	stmt, err := connection.PrepareContext(ctx, `
		select users.user_id, users.language, user.prefixes
		from users
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
