package directives

import (
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"
	"github.com/PrameshKarki/event-management-golang/utils"
)

func Private() func(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	return func(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
		// Get user if from context
		user := ctx.Value("user")
		if user != nil {
			userID := user.(*utils.TokenMetadata).ID
			if userID == "" {
				return nil, errors.New("unauthorized")

			} else {
				return next(ctx)
			}
		} else {
			return nil, errors.New("unauthorized")
		}
	}
}
