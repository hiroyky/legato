package dto

import "github.com/volatiletech/sqlboiler/v4/queries/qm"

type GetTracksDTO struct {
	TrackID       *int
	AlbumArtistID *int
	AlbumID       *int
	GenreID       *int
	Limit         *int
	Offset        *int
}

func (d *GetTracksDTO) GenWhereMods() []qm.QueryMod {
	mods := make([]qm.QueryMod, 0)

	if d == nil {
		return mods
	}

	if d.TrackID != nil {
		mods = append(mods, qm.Where("track_id=?", d.TrackID))
	}
	if d.AlbumArtistID != nil {
		mods = append(mods, qm.Where("album_artist_id=?", d.TrackID))
	}
	if d.AlbumID != nil {
		mods = append(mods, qm.Where("album_id=?", d.TrackID))
	}
	if d.GenreID != nil {
		mods = append(mods, qm.Where("genre_id=?", d.TrackID))
	}

	return mods
}
