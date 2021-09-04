package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/legato/infrastructure/database/repository/dto"

	"github.com/legato/graph/generated"
	"github.com/legato/graph/gqlmodel"
	"github.com/legato/lib/gql"
)

func (r *albumResolver) AlbumArtist(ctx context.Context, obj *gqlmodel.Album) (*gqlmodel.AlbumArtist, error) {
	albumArtistID, err := gql.DecodeID(obj.AlbumArtistID)
	if err != nil {
		return nil, err
	}
	albumArtist, err := r.AlbumArtistRepository.GetByID(ctx, albumArtistID)
	if err != nil {
		return nil, err
	}
	return gqlmodel.NewAlbumArtist(albumArtist), nil
}

func (r *albumResolver) Tracks(ctx context.Context, obj *gqlmodel.Album) ([]*gqlmodel.Track, error) {
	albumID, err := gql.DecodedIDIntPtr(obj.ID)
	data := &dto.GetTracksDTO{
		AlbumID: albumID,
	}
	tracks, err := r.TrackRepository.GetTracks(ctx, data)
	if err != nil {
		return nil, err
	}
	return gqlmodel.NewTracks(tracks), err
}

// Album returns generated.AlbumResolver implementation.
func (r *Resolver) Album() generated.AlbumResolver { return &albumResolver{r} }

type albumResolver struct{ *Resolver }
