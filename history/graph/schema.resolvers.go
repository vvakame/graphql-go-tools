package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"gist.github.com/vvakame/7a08b7f517b012cca09f82ec9ab4d1fd/history/graph/generated"
	"gist.github.com/vvakame/7a08b7f517b012cca09f82ec9ab4d1fd/history/graph/model"
)

func (r *userResolver) Histories(ctx context.Context, obj *model.User, first *int, after *string) (*model.HistoryConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
