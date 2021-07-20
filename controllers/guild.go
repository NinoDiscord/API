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

func (c *Controller) GetGuild(
	context context.Context,
	conn *sql.DB,
	id string,
) (*types.Guild, error) {
	guilds, err := c.GetGuilds(context, conn); if err != nil {
		return nil, err
	}

	var guild *types.Guild
	for _, g := range guilds {
		if g.ID == id {
			guild = g
			break
		}
	}

	return guild, nil
}

func (c *Controller) AddGuildPrefix(
	context context.Context,
	conn *sql.DB,
	id string,
	prefix string,
) (bool, error) {
	guild, err := c.GetGuild(context, conn, id); if err != nil {
		return false, err
	}

	if GuildPrefixExists(guild, prefix) {
		return false, errors.New(fmt.Sprintf("prefix %s already exists in the db", prefix))
	}

	stmt, err := conn.PrepareContext(context, `
		UPDATE guilds SET prefixes = array_append(prefixes, $1) WHERE guild_id = $2
	`); if err != nil {
		return false, err
	}

	_, err = stmt.Query(prefix, guild.ID)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (c *Controller) RemoveGuildPrefix(
	context context.Context,
	conn *sql.DB,
	id string,
	prefix string,
) (bool, error) {
	guild, err := c.GetGuild(context, conn, id);
	if err != nil {
		return false, err
	}

	if !GuildPrefixExists(guild, prefix) {
		return false, errors.New(fmt.Sprintf("prefix %s was not a  prefix", prefix))
	}

	stmt, err := conn.PrepareContext(context, `
		UPDATE guilds SET prefixes = array_remove(prefixes, $1) WHERE guild_id = $2
	`); if err != nil {
		return false, err
	}

	_, err = stmt.Query(prefix, guild.ID)
	if err != nil {
		return false, err
	}

	return true, nil
}

func GuildPrefixExists(guild *types.Guild, prefix string) bool {
	for _, pre := range guild.Prefixes {
		if prefix == pre {
			return true
		}
	}

	return false
}
