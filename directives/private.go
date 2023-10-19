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
		if user == nil {
			return nil, errors.New("unauthenticated")
		}
		if user.(*utils.TokenMetadata).ID != "" {
			return next(ctx)
		}
		return nil, errors.New("unauthenticated")

	}
}
