package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"gist.github.com/vvakame/7a08b7f517b012cca09f82ec9ab4d1fd/history/graph/generated"
	"gist.github.com/vvakame/7a08b7f517b012cca09f82ec9ab4d1fd/history/graph/model"
)

func strToPtr(s string) *string {
	return &s
}

func (r *userResolver) Histories(ctx context.Context, obj *model.User, first *int, after *string) (*model.HistoryConnection, error) {
	histories := []*model.History{
		{
			ID: "1",
			Detail: &model.HistoryEntityA{
				ID:    strToPtr("1a"),
				TextA: strToPtr("some text 1"),
			},
		},
		{
			ID: "2",
			Detail: &model.HistoryEntityB{
				ID:    strToPtr("2b"),
				TextB: strToPtr("some text 2"),
			},
		},
		{
			ID:     "3",
			Detail: nil,
		},
	}

	edges := make([]*model.HistoryEdge, len(histories))

	for i := range histories {
		edges[i] = &model.HistoryEdge{
			Node:   histories[i],
			Cursor: &histories[i].ID,
		}
	}

	return &model.HistoryConnection{
		Nodes: histories,
		Edges: edges,
	}, nil
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
