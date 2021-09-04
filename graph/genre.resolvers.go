package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/legato/graph/generated"
	"github.com/legato/graph/gqlmodel"
	"github.com/legato/infrastructure/database/repository/dto"
	"github.com/legato/lib/gql"
)

func (r *genreResolver) TrackPagination(ctx context.Context, obj *gqlmodel.Genre, limit int, offset *int) (*gqlmodel.TrackPagination, error) {
	genreID, err := gql.DecodedIDIntPtr(obj.ID)
	if err != nil {
		return nil, err
	}
	data := &dto.GetTracksDTO{
		GenreID: genreID,
		Limit:   &limit,
		Offset:  offset,
	}
	tracks, err := r.TrackRepository.GetTracks(ctx, data)
	if err != nil {
		return nil, err
	}
	total, err := r.TrackRepository.CountTracks(ctx, data)
	if err != nil {
		return nil, err
	}
	return gqlmodel.NewTrackPagination(tracks, total, limit, offset), nil
}

// Genre returns generated.GenreResolver implementation.
func (r *Resolver) Genre() generated.GenreResolver { return &genreResolver{r} }

type genreResolver struct{ *Resolver }
