package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/jjjjackson/gqlgen-example/domain/model"
	"github.com/jjjjackson/gqlgen-example/graphql/generated"
	"github.com/jjjjackson/gqlgen-example/infrastructure/datastore"
	"github.com/jjjjackson/gqlgen-example/presentation/iotype"
	"github.com/jjjjackson/gqlgen-example/usecase"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input iotype.CreateUserInput) (*model.User, error) {
	userService := usecase.NewUserService(
		r.DB,
		datastore.NewUserRepository(),
	)
	return userService.CreateUser(input.Email, input.Password)
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	userService := usecase.NewUserService(
		r.DB,
		datastore.NewUserRepository(),
	)
	return userService.ListUsers()
}

func (r *queryResolver) User(ctx context.Context, email string) (*model.User, error) {
	userService := usecase.NewUserService(
		r.DB,
		datastore.NewUserRepository(),
	)
	return userService.GetUser(email)
}

func (r *userResolver) CreatedAt(ctx context.Context, obj *model.User) (string, error) {
	return obj.CreatedAt.String(), nil
}

func (r *userResolver) UpdatedAt(ctx context.Context, obj *model.User) (string, error) {
	return obj.UpdatedAt.String(), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
