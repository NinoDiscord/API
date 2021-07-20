package controllers

import (
	"context"
	"database/sql"
	"github.com/lib/pq"
	"nino.sh/api/graphql/types"
	"nino.sh/api/utils"
)

// GetGuilds returns an array of types.Guild objects
func (c *Controller) GetGuilds(
	context context.Context,
	connection *sql.DB,
) ([]*types.Guild, error) {
	stmt, err := connection.PrepareContext(context, `
		select guilds.guild_id, guilds.modlog_channel_id, guilds.muted_role_id, guilds.prefixes, guilds.language
		from guilds
	`); if err != nil {
		return nil, err
	}

	var guilds []*types.Guild
	rows, err := stmt.QueryContext(context)
	if err != nil {
		return nil, err
	}

	defer utils.SwallowError(rows)
	for rows.Next() {
		// i cant believe i have to do this
		var (
			modLogID *string
			mutedRoleID *string
			prefixes []string
			language string
			id string
		)

		err = rows.Scan(&id, &modLogID, &mutedRoleID, pq.Array(&prefixes), &language)
		if err != nil {
			return nil, err
		}

		guilds = append(guilds, &types.Guild{
			ModLogChannelID: modLogID,
			MutedRoleID: mutedRoleID,
			Prefixes: prefixes,
			Language: language,
			ID: id,
		})
	}

	return guilds, nil
}
