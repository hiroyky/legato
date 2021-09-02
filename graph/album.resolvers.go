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
	panic(fmt.Errorf("not implemented"))
}

// Album returns generated.AlbumResolver implementation.
func (r *Resolver) Album() generated.AlbumResolver { return &albumResolver{r} }

type albumResolver struct{ *Resolver }
