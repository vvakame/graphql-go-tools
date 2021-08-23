package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"gist.github.com/vvakame/7a08b7f517b012cca09f82ec9ab4d1fd/base/graph/generated"
	"gist.github.com/vvakame/7a08b7f517b012cca09f82ec9ab4d1fd/base/graph/model"
)

func (r *queryResolver) Me(ctx context.Context) (*model.User, error) {
	return &model.User{
		ID: "1234",
	}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
