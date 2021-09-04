package dto

import "github.com/volatiletech/sqlboiler/v4/queries/qm"

type GetAlbumsDTO struct {
	Limit         *int
	Offset        *int
	AlbumID       *int
	AlbumArtistID *int
}

func (d *GetAlbumsDTO) GenWhereMods() []qm.QueryMod {
	mods := make([]qm.QueryMod, 0)

	if d == nil {
		return mods
	}

	if d.AlbumID != nil {
		mods = append(mods, qm.Where("album_id=?", d.AlbumID))
	}
	if d.AlbumArtistID != nil {
		mods = append(mods, qm.Where("album_artist_id=?", d.AlbumArtistID))
	}

	return mods
}
