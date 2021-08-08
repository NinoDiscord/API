package controllers

import (
	"context"
	"database/sql"
	"errors"
	"nino.sh/api/graphql/types"
	"nino.sh/api/utils"
)

type AutomodController struct {}

// GetAll returns an array of types.Automod objects.
func (c *AutomodController) GetAll(
	ctx context.Context,
	conn *sql.DB,
) ([]*types.Automod, error) {
	stmt, err := conn.PrepareContext(ctx, `
		select automod.blacklist, automod.mentions, automod.invites, automod.dehoist,
		automod.guild_id, automod.spam, automod.raid, automod.blacklist_words, automod.short_links,
		automod.whitelist_channels_during_raid from automod
	`); if err != nil {
		return nil, err
	}

	var automod []*types.Automod
	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, err
	}

	defer utils.SwallowError(rows)
	for rows.Next() {
		var (
			blacklist bool
			mentions bool
			invites bool
			dehoist bool
			guildId string
			spam bool
			raid bool
			blacklistWords []string
			shortLinks bool
			whitelistChannels []string
		)

		err = rows.Scan(
			&blacklist,
			&mentions,
			&invites,
			&dehoist,
			&guildId,
			&spam,
			&raid,
			&blacklistWords,
			&shortLinks,
			&whitelistChannels,
		)

		if err != nil {
			return nil, err
		}

		automod = append(automod, &types.Automod{
			WhitelistChannels: whitelistChannels,
			BlacklistedWords: blacklistWords,
			ShortLinks: shortLinks,
			Blacklist: blacklist,
			Mentions: mentions,
			GuildId: guildId,
			Invites: invites,
			Dehoist: dehoist,
			Spam: spam,
			Raid: raid,
		})
	}

	return automod, nil
}

func (c *AutomodController) Get(
	ctx context.Context,
	conn *sql.DB,
	id string,
) (*types.Automod, error) {
	objects, err := c.GetAll(ctx, conn); if err != nil {
		return nil, err
	}

	if id == "" {
		return nil, errors.New("id must be specified")
	}

	var automod *types.Automod
	for _, auto := range objects {
		if auto.GuildId == id {
			automod = auto
			break
		}
	}

	return automod, nil
}
