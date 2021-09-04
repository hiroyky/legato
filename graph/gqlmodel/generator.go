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
		ID:       gql.EncodeID(TrackName, track.TrackID),
		Title:    track.Title,
		Artist:   track.Artist,
		Composer: track.Composer,
		TrackNo:  track.TrackNo,
		Lyrics:   track.Lyrics,
		Comment:  track.Comment,
		Year:     track.Year,
		URL: fmt.Sprintf(
			"%s://%s:%d/music/%s",
			config.Env.HTTPProtocol,
			config.Env.APIHostName,
			config.Env.APIPort,
			track.FilePathHash,
		),
		DownloadURL: fmt.Sprintf(
			"%s://%s:%d/music/%s/download",
			config.Env.HTTPProtocol,
			config.Env.APIHostName,
			config.Env.APIPort,
			track.FilePathHash,
		),
		AlbumID:       gql.EncodeID(AlbumName, track.AlbumID),
		GenreID:       gql.EncodeID(GenreName, track.GenreID),
		AlbumArtistID: gql.EncodeID(AlbumArtistName, track.AlbumArtistID),
	}
}

func NewAlbumPagination(albums dbmodel.AlbumSlice, total int64, limit int, offset *int) *AlbumPagination {
	return &AlbumPagination{
		PageInfo: newPaginationInfo(total, len(albums), limit, offset),
		Edges:    newAlbumEdges(albums),
		Nodes:    NewAlbums(albums),
	}
}

func newAlbumEdges(albums dbmodel.AlbumSlice) []*AlbumEdge {
	edges := make([]*AlbumEdge, len(albums))
	for i, album := range albums {
		edges[i] = &AlbumEdge{
			Cursor: gql.EncodeID(AlbumName, album.AlbumID),
			Node:   NewAlbum(album),
		}
	}
	return edges
}

func NewAlbums(albums dbmodel.AlbumSlice) []*Album {
	slice := make([]*Album, len(albums))
	for i, v := range albums {
		slice[i] = NewAlbum(v)
	}
	return slice
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

func NewAlbumArtistPagination(albumArtists dbmodel.AlbumArtistSlice, total int64, limit int, offset *int) *AlbumArtistPagination {
	return &AlbumArtistPagination{
		PageInfo: newPaginationInfo(total, len(albumArtists), limit, offset),
		Edges:    newAlbumArtistEdges(albumArtists),
		Nodes:    NewAlbumArtists(albumArtists),
	}
}

func newAlbumArtistEdges(albumArtists dbmodel.AlbumArtistSlice) []*AlbumArtistEdge {
	edges := make([]*AlbumArtistEdge, len(albumArtists))
	for i, aa := range albumArtists {
		edges[i] = &AlbumArtistEdge{
			Cursor: gql.EncodeID(AlbumArtistName, aa.AlbumArtistID),
			Node:   NewAlbumArtist(aa),
		}
	}
	return edges
}

func NewAlbumArtists(albumArtists dbmodel.AlbumArtistSlice) []*AlbumArtist {
	slice := make([]*AlbumArtist, len(albumArtists))
	for i, v := range albumArtists {
		slice[i] = NewAlbumArtist(v)
	}
	return slice
}

func NewAlbumArtist(albumArtist *dbmodel.AlbumArtist) *AlbumArtist {
	return &AlbumArtist{
		ID:   gql.EncodeID(AlbumArtistName, albumArtist.AlbumArtistID),
		Name: albumArtist.Name,
	}
}

func NewGenrePagination(genres dbmodel.GenreSlice, total int64, limit int, offset *int) *GenrePagination {
	return &GenrePagination{
		PageInfo: newPaginationInfo(total, len(genres), limit, offset),
		Edges:    newGenreEdges(genres),
		Nodes:    NewGenres(genres),
	}
}

func newGenreEdges(genres dbmodel.GenreSlice) []*GenreEdge {
	edges := make([]*GenreEdge, len(genres))
	for i, genre := range genres {
		edges[i] = &GenreEdge{
			Cursor: gql.EncodeID(GenreName, genre.GenreID),
			Node:   NewGenre(genre),
		}
	}
	return edges
}

func NewGenres(genres dbmodel.GenreSlice) []*Genre {
	slice := make([]*Genre, len(genres))
	for i, v := range genres {
		slice[i] = NewGenre(v)
	}
	return slice
}

func NewGenre(genre *dbmodel.Genre) *Genre {
	return &Genre{
		ID:   gql.EncodeID(GenreName, genre.GenreID),
		Name: genre.Name,
	}
}
