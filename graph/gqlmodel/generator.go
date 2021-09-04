package gqlmodel

import (
	"fmt"
	"github.com/legato/infrastructure/config"
	"github.com/legato/infrastructure/database/dbmodel"
	"github.com/legato/lib/gql"
)

func newPaginationInfo(total int64, count, limit int, offset *int) *PaginationInfo {
	o := 0
	t := int(total)
	if offset != nil {
		o = *offset
	}
	page := 1
	pageLength := 1
	if limit > 0 {
		page = (o / limit) + 1
		pageLength = (t / limit) + 1
	}

	return &PaginationInfo{
		Page:             page,
		PaginationLength: pageLength,
		HasNextPage:      o+limit < t,
		HasPreviousPage:  o > 0,
		Count:            count,
		TotalCount:       t,
		Limit:            limit,
		Offset:           o,
	}
}

func NewTrackPagination(tracks dbmodel.TrackSlice, total int64, limit int, offset *int) *TrackPagination {
	return &TrackPagination{
		PageInfo: newPaginationInfo(total, len(tracks), limit, offset),
		Edges:    newTrackEdges(tracks),
		Nodes:    NewTracks(tracks),
	}
}

func newTrackEdges(tracks dbmodel.TrackSlice) []*TrackEdge {
	edges := make([]*TrackEdge, len(tracks))
	for i, track := range tracks {
		edges[i] = &TrackEdge{
			Cursor: gql.EncodeID(TrackName, track.TrackID),
			Node:   NewTrack(track),
		}
	}
	return edges
}

func NewTracks(tracks dbmodel.TrackSlice) []*Track {
	slice := make([]*Track, len(tracks))
	for i, v := range tracks {
		slice[i] = NewTrack(v)
	}
	return slice
}

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

func NewGenre(genre *dbmodel.Genre) *Genre {
	return &Genre{
		ID:   gql.EncodeID(GenreName, genre.GenreID),
		Name: genre.Name,
	}
}
