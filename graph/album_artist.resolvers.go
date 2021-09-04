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

func (r *albumArtistResolver) Albums(ctx context.Context, obj *gqlmodel.AlbumArtist) ([]*gqlmodel.Album, error) {
	albumArtistID, err := gql.DecodedIDIntPtr(obj.ID)
	if err != nil {
		return nil, err
	}
	data := &dto.GetAlbumsDTO{
		AlbumArtistID: albumArtistID,
	}
	albums, err := r.AlbumRepository.GetAlbums(ctx, data)
	if err != nil {
		return nil, err
	}
	return gqlmodel.NewAlbums(albums), nil
}

func (r *albumArtistResolver) AlbumPagination(ctx context.Context, obj *gqlmodel.AlbumArtist, limit int, offset *int) (*gqlmodel.AlbumPagination, error) {
	albumArtistID, err := gql.DecodedIDIntPtr(obj.ID)
	if err != nil {
		return nil, err
	}
	data := &dto.GetAlbumsDTO{
		AlbumArtistID: albumArtistID,
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

func (r *albumArtistResolver) TrackPagination(ctx context.Context, obj *gqlmodel.AlbumArtist, limit int, offset *int) (*gqlmodel.TrackPagination, error) {
	albumArtistID, err := gql.DecodedIDIntPtr(obj.ID)
	if err != nil {
		return nil, err
	}
	data := &dto.GetTracksDTO{
		AlbumArtistID: albumArtistID,
		Limit:         &limit,
		Offset:        offset,
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

// AlbumArtist returns generated.AlbumArtistResolver implementation.
func (r *Resolver) AlbumArtist() generated.AlbumArtistResolver { return &albumArtistResolver{r} }

type albumArtistResolver struct{ *Resolver }
