package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/hizzuu/beatic-backend/graph/generated"
	"github.com/hizzuu/beatic-backend/graph/model"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*model.CreateUserPayload, error) {
	panic(fmt.Errorf("not implemented: CreateUser - createUser"))
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UpdateUserInput) (*model.UpdateUserPayload, error) {
	panic(fmt.Errorf("not implemented: UpdateUser - updateUser"))
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id int64) (*model.GetUserPayload, error) {
	r.logger.Infof(ctx, "get by user query")
	return nil, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
