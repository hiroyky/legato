package gqlmodel

import (
	"fmt"
	"github.com/legato/infrastructure/config"
	"github.com/legato/infrastructure/database/dbmodel"
	"github.com/legato/lib/gql"
)

func NewTrack(track *dbmodel.Track) *Track {
	return &Track{
		ID:            gql.EncodeID(TrackName, track.TrackID),
		Title:         track.Title,
		Artist:        track.Artist,
		Composer:      track.Composer,
		TrackNo:       track.TrackNo,
		Lyrics:        track.Lyrics,
		Comment:       track.Comment,
		Year:          track.Year,
		URL:           fmt.Sprintf("%s://%s:%d/music/%s", config.Env.HTTPProtocol, config.Env.APIHostName, config.Env.APIPort, track.FilePathHash),
		AlbumID:       gql.EncodeID(AlbumName, track.AlbumID),
		GenreID:       gql.EncodeID(GenreName, track.GenreID),
		AlbumArtistID: gql.EncodeID(AlbumArtistName, track.AlbumArtistID),
	}
}

func NewAlbum(album *dbmodel.Album) *Album {
	return &Album{
		ID:            gql.EncodeID(AlbumName, album.AlbumID),
		Name:          album.Name,
		DiskNo:        album.DiscNo,
		DiskTotal:     album.DiscTotal,
		AlbumArtistID: gql.EncodeID(AlbumArtistName, album.AlbumArtistID),
	}
}

func NewAlbumArtist(albumArtist *dbmodel.AlbumArtist) *AlbumArtist {
	return &AlbumArtist{
		ID:   gql.EncodeID(AlbumArtistName, albumArtist.AlbumArtistID),
		Name: albumArtist.Name,
	}
}
