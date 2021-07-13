package resolvers

import (
	"context"
	"database/sql"
	"github.com/sirupsen/logrus"
	"nino.sh/api/graphql/types"
)

// Guild queries a guild's metadata by it's ID, authentication is required on this query.
func (r *Resolver) Guild(ctx context.Context, args struct{ ID string }) (*types.Guild, error) {
	stmt, err := r.Db.Connection.PrepareContext(ctx, "SELECT * FROM guilds WHERE guild_id = $1")
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	logrus.Info()

	var guild *types.Guild
	err = stmt.QueryRowContext(ctx, args.ID).Scan(&guild)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return guild, nil
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
