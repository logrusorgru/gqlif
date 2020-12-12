package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/logrusorgru/gqlif/internal/graph/server"
	"github.com/logrusorgru/gqlif/internal/inputs"
	"github.com/logrusorgru/gqlif/internal/types"
)

func (r *mutationResolver) Set(ctx context.Context, name string, value inputs.InputValue) (*types.Value, error) {
	return r.Resolver.Set(ctx, name, value)
}

func (r *queryResolver) Get(ctx context.Context, name string) (*types.Value, error) {
	return r.Resolver.Get(ctx, name)
}

// Mutation returns server.MutationResolver implementation.
func (r *Resolver) Mutation() server.MutationResolver { return &mutationResolver{r} }

// Query returns server.QueryResolver implementation.
func (r *Resolver) Query() server.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
