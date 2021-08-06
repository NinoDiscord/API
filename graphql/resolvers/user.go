package resolvers

import (
	"context"
	"nino.sh/api/graphql/types"
)

// User queries a user's metadata by its ID, authentication is not required on this query.
func (r *Resolver) User(ctx context.Context, args struct { ID string }) (*types.User, error) {
	user, err := r.Controller.GetUser(ctx, r.Db.Connection, args.ID); if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *Resolver) AddUserPrefix(ctx context.Context, args struct { ID string; Prefix string }) (bool, error) {
	if err := r.CheckAuthorization(ctx.Value("token").(string)); err != nil {
		return false, err
	}

	return r.Controller.AddUserPrefix(ctx, r.Db.Connection, args.ID, args.Prefix)
}

func (r *Resolver) RemoveUserPrefix(ctx context.Context, args struct { ID string; Prefix string }) (bool, error) {
	if err := r.CheckAuthorization(ctx.Value("token").(string)); err != nil {
		return false, err
	}

	return r.Controller.RemoveUserPrefix(ctx, r.Db.Connection, args.ID, args.Prefix)
}

func (r *Resolver) UpdateUserLanguage(ctx context.Context, args struct { ID string; Language string }) (bool, error) {
	if err := r.CheckAuthorization(ctx.Value("token").(string)); err != nil {
		return false, err
	}

	return r.Controller.UpdateUserLanguage(ctx, r.Db.Connection, args.ID, args.Language)
}
