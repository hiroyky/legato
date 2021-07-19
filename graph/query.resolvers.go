package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/legato/graph/generated"
	"github.com/legato/graph/model"
)

func (r *queryResolver) Track(ctx context.Context, id string) (*model.Track, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Tracks(ctx context.Context, limit int, offset *int) (*model.TrackPagination, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Album(ctx context.Context, id string) (*model.Album, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Albums(ctx context.Context, limit int, offset *int) (*model.AlbumPagination, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AlbumArtist(ctx context.Context, id string) (*model.AlbumArtist, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AlbumArtists(ctx context.Context, limit int, offset *int) (*model.AlbumArtistPagination, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Genre(ctx context.Context, id string) (*model.Genre, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Genres(ctx context.Context, limit int, offset *int) (*model.GenrePagination, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
