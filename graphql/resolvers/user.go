package resolvers

import (
	"context"
	"nino.sh/api/graphql/types"
)

// User queries a user's metadata by it's ID, authentication is required on this query.
func (r *Resolver) User(ctx context.Context, args struct { ID string }) (*types.User, error) {
	row := r.Db.Connection.QueryRowContext(ctx, "SELECT * FROM guilds WHERE guild_id = ?", args.ID)

	var user *types.User
	if err := row.Scan(&user); err != nil {
		return nil, err
	}

	return user, nil
}

/*
func (r *Resolver) UpdateUser(id string) (string, error) {
	return id, nil
}

func (r *Resolver) RemoveUserPrefix(id string) (string, error) {
	return id, nil
}

func (r *Resolver) AddUserPrefix(id string) (string, error) {
	return id, nil
}
*/
