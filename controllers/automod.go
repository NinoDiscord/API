package controllers

import (
	"context"
	"database/sql"
	"nino.sh/api/graphql/types"
)

// GetAutomod returns the automod object of a guild, or nil
// if the guild doesn't have Nino in it
func (c *Controller) GetAutomod(
	ctx context.Context,
	conn *sql.Conn,
	id string,
) (*types.Automod, error) {
	return nil, nil
}

/*
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
 */
