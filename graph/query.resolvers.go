package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/legato/graph/generated"
	"github.com/legato/graph/gqlmodel"
	"github.com/legato/lib/gql"
)

func (r *queryResolver) Track(ctx context.Context, id string) (*gqlmodel.Track, error) {
	decodedID, err := gql.DecodeID(id)
	if err != nil {
		return nil, err
	}

	track, err := r.TrackRepository.GetByID(ctx, decodedID)
	if err != nil {
		return nil, err
	}

	return gqlmodel.NewTrack(track), nil
}

func (r *queryResolver) Tracks(ctx context.Context, limit int, offset *int) (*gqlmodel.TrackPagination, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Album(ctx context.Context, id string) (*gqlmodel.Album, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Albums(ctx context.Context, limit int, offset *int) (*gqlmodel.AlbumPagination, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AlbumArtist(ctx context.Context, id string) (*gqlmodel.AlbumArtist, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AlbumArtists(ctx context.Context, limit int, offset *int) (*gqlmodel.AlbumArtistPagination, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Genre(ctx context.Context, id string) (*gqlmodel.Genre, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Genres(ctx context.Context, limit int, offset *int) (*gqlmodel.GenrePagination, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
