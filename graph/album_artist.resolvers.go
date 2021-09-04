package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/legato/graph/generated"
	"github.com/legato/graph/gqlmodel"
)

func (r *albumArtistResolver) Albums(ctx context.Context, obj *gqlmodel.AlbumArtist) ([]*gqlmodel.Album, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *albumArtistResolver) AlbumPagination(ctx context.Context, obj *gqlmodel.AlbumArtist, limit int, offset *int) (*gqlmodel.AlbumArtistPagination, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *albumArtistResolver) TrackPagination(ctx context.Context, obj *gqlmodel.AlbumArtist, limit int, offset *int) (*gqlmodel.TrackPagination, error) {
	panic(fmt.Errorf("not implemented"))
}

// AlbumArtist returns generated.AlbumArtistResolver implementation.
func (r *Resolver) AlbumArtist() generated.AlbumArtistResolver { return &albumArtistResolver{r} }

type albumArtistResolver struct{ *Resolver }
