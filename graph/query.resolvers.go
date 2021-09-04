package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/legato/graph/generated"
	"github.com/legato/graph/gqlmodel"
	"github.com/legato/infrastructure/database/repository/dto"
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
	data := &dto.GetTracksDTO{
		Limit:  &limit,
		Offset: offset,
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

func (r *queryResolver) Album(ctx context.Context, id string) (*gqlmodel.Album, error) {
	decodedID, err := gql.DecodeID(id)
	if err != nil {
		return nil, err
	}
	album, err := r.AlbumRepository.GetByID(ctx, decodedID)
	if err != nil {
		return nil, err
	}
	return gqlmodel.NewAlbum(album), nil
}

func (r *queryResolver) Albums(ctx context.Context, limit int, offset *int) (*gqlmodel.AlbumPagination, error) {
	data := &dto.GetAlbumsDTO{
		Limit:  &limit,
		Offset: offset,
	}
	albums, err := r.AlbumRepository.GetAlbums(ctx, data)
	if err != nil {
		return nil, err
	}
	total, err := r.AlbumRepository.CountAlbums(ctx, data)
	if err != nil {
		return nil, err
	}
	return gqlmodel.NewAlbumPagination(albums, total, limit, offset), nil
}

func (r *queryResolver) AlbumArtist(ctx context.Context, id string) (*gqlmodel.AlbumArtist, error) {
	decodedID, err := gql.DecodeID(id)
	if err != nil {
		return nil, err
	}
	albumArtist, err := r.AlbumArtistRepository.GetByID(ctx, decodedID)
	if err != nil {
		return nil, err
	}
	return gqlmodel.NewAlbumArtist(albumArtist), nil
}

func (r *queryResolver) AlbumArtists(ctx context.Context, limit int, offset *int) (*gqlmodel.AlbumArtistPagination, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Genre(ctx context.Context, id string) (*gqlmodel.Genre, error) {
	decodedID, err := gql.DecodeID(id)
	if err != nil {
		return nil, err
	}
	genre, err := r.GenreRepository.GetByID(ctx, decodedID)
	if err != nil {
		return nil, err
	}
	return gqlmodel.NewGenre(genre), nil
}

func (r *queryResolver) Genres(ctx context.Context, limit int, offset *int) (*gqlmodel.GenrePagination, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
