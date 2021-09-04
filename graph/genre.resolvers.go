package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/legato/graph/generated"
	"github.com/legato/graph/gqlmodel"
)

func (r *genreResolver) TrackPagination(ctx context.Context, obj *gqlmodel.Genre, limit int, offset *int) (*gqlmodel.TrackPagination, error) {
	panic(fmt.Errorf("not implemented"))
}

// Genre returns generated.GenreResolver implementation.
func (r *Resolver) Genre() generated.GenreResolver { return &genreResolver{r} }

type genreResolver struct{ *Resolver }
