package resolvers

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/sirupsen/logrus"
	"nino.sh/api/graphql/types"
)

// Guild queries a guild's metadata by it's ID, authentication is required on this query.
func (r *Resolver) Guild(ctx context.Context, args struct{ ID string }) (*types.Guild, error) {
	row := r.Db.Connection.QueryRowContext(ctx, fmt.Sprintf("SELECT * FROM guilds WHERE guild_id = %s;", args.ID))

	var guild *types.Guild
	err := row.Scan(&guild)

	switch {
	case err == sql.ErrNoRows:
		return nil, nil

	case err != nil:
		logrus.Error(err.Error())
		return nil, err

	default:
		return guild, nil
	}
}

/*
func (r *Resolver) UpdateGuild(id string) (string, error) {
	return id, nil
}

func (r *Resolver) RemoveGuildPrefix(id string) (string, error) {
	return id, nil
}

func (r *Resolver) AddGuildPrefix(id string) (string, error) {
	return id, nil
}
*/
