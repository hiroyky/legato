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

func (r *trackResolver) Album(ctx context.Context, obj *gqlmodel.Track) (*gqlmodel.Album, error) {
	decodedID, err := gql.DecodeID(obj.AlbumID)
	if err != nil {
		return nil, err
	}
	album, err := r.AlbumRepository.GetByID(ctx, decodedID)
	if err != nil {
		return nil, err
	}

	return gqlmodel.NewAlbum(album), nil
}

func (r *trackResolver) Genre(ctx context.Context, obj *gqlmodel.Track) (*gqlmodel.Genre, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *trackResolver) AlbumArtist(ctx context.Context, obj *gqlmodel.Track) (*gqlmodel.AlbumArtist, error) {
	panic(fmt.Errorf("not implemented"))
}

// Track returns generated.TrackResolver implementation.
func (r *Resolver) Track() generated.TrackResolver { return &trackResolver{r} }

type trackResolver struct{ *Resolver }
