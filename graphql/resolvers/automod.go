package resolvers

import (
	"context"
	"nino.sh/api/graphql/types"
)

func (r *Resolver) Automod(ctx context.Context, args struct { ID string }) (*types.Automod, error) {
	if err := r.CheckAuthorization(ctx.Value("token").(string)); err != nil {
		return nil, err
	}

	return r.Controller.Automod.Get(ctx, r.Db.Connection, args.ID)
}

/*
func (r *Resolver) ToggleShortlinks(id string) (string, error) {
	return id, nil
}

func (r *Resolver) ToggleBlacklist(id string) (string, error) {
	return id, nil
}

func (r *Resolver) ToggleMentions(id string) (string, error) {
	return id, nil
}

func (r *Resolver) ToggleInvites(id string) (string, error) {
	return id, nil
}

func (r *Resolver) ToggleDehoist(id string) (string, error) {
	return id, nil
}

func (r *Resolver) ToggleSpam(id string) (string, error) {
	return id, nil
}

func (r *Resolver) ToggleRaid(id string) (string, error) {
	return id, nil
}
*/
