package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/hiroyky/legato/graph/generated"
	"github.com/hiroyky/legato/graph/gqlmodel"
	"github.com/hiroyky/legato/lib/gql"
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
	decodedID, err := gql.DecodeID(obj.GenreID)
	if err != nil {
		return nil, err
	}
	genre, err := r.GenreRepository.GetByID(ctx, decodedID)
	if err != nil {
		return nil, err
	}
	return gqlmodel.NewGenre(genre), nil
}

func (r *trackResolver) AlbumArtist(ctx context.Context, obj *gqlmodel.Track) (*gqlmodel.AlbumArtist, error) {
	decodedID, err := gql.DecodeID(obj.AlbumArtistID)
	if err != nil {
		return nil, err
	}
	albumArtist, err := r.AlbumArtistRepository.GetByID(ctx, decodedID)
	if err != nil {
		return nil, err
	}
	return gqlmodel.NewAlbumArtist(albumArtist), nil
}

// Track returns generated.TrackResolver implementation.
func (r *Resolver) Track() generated.TrackResolver { return &trackResolver{r} }

type trackResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *trackResolver) DownURL(ctx context.Context, obj *gqlmodel.Track) (string, error) {
	panic(fmt.Errorf("not implemented"))
}
