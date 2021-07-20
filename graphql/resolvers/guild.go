package resolvers

import (
	"context"
	"nino.sh/api/graphql/types"
)

// Guild queries a guild's metadata by it's ID, authentication is required on this query.
func (r *Resolver) Guild(ctx context.Context, args struct{ ID string }) (*types.Guild, error) {
	guilds, err := r.Controller.GetGuilds(ctx, r.Db.Connection); if err != nil {
		return nil, err
	}

	var guild *types.Guild
	for _, g := range guilds {
		if g.ID == args.ID {
			guild = g
			break
		}
	}

	if guild != nil {
		return guild, nil
	} else {
		return nil, nil
	}
}
