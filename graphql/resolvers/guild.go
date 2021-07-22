package resolvers

import (
	"context"
	"nino.sh/api/graphql/types"
)

// Guild queries a guild's metadata by it's ID, authentication is not required on this query (since it doesn't have sensitive data).
func (r *Resolver) Guild(ctx context.Context, args struct{ ID string }) (*types.Guild, error) {
	return r.Controller.GetGuild(ctx, r.Db.Connection, args.ID)
}

// AddGuildPrefix is a mutation to add a prefix to the database, authentication is required on this query.
func (r *Resolver) AddGuildPrefix(ctx context.Context, args struct{ ID string; Prefix string }) (bool, error) {
	return r.Controller.AddGuildPrefix(ctx, r.Db.Connection, args.ID, args.Prefix)
}

// RemoveGuildPrefix is a mutation to remove a prefix from the database, authentication is required on this query.
func (r *Resolver) RemoveGuildPrefix(ctx context.Context, args struct{ ID string; Prefix string }) (bool, error) {
	return r.Controller.RemoveGuildPrefix(ctx, r.Db.Connection, args.ID, args.Prefix)
}

// UpdateModLog is a mutation to denote a mod-log channel
func (r *Resolver) UpdateModLog(ctx context.Context, args struct{ ID string; ModLogId *string }) (bool, error) {
	return r.Controller.UpdateMutedRole(ctx, r.Db.Connection, args.ID, args.ModLogId)
}

// UpdateMutedRole is a mutation to denote a Muted role
func (r *Resolver) UpdateMutedRole(ctx context.Context, args struct{ ID string; MutedRoleID *string }) (bool, error) {
	return r.Controller.UpdateMutedRole(ctx, r.Db.Connection, args.ID, args.MutedRoleID)
}
