package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.39

import (
	"context"
	"fmt"

	"github.com/PrameshKarki/event-management-golang/graph/model"
	services "github.com/PrameshKarki/event-management-golang/graph/services/auth"
)

// UserSignUp is the resolver for the userSignUp field.
func (r *mutationResolver) UserSignUp(ctx context.Context, data model.UserInput) (*model.AuthSchema, error) {
	userId, _ := services.UserSignUp(data)
	return &model.AuthSchema{ID: fmt.Sprint(userId), AccessToken: "Test"}, nil
}