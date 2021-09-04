package dto

import "github.com/volatiletech/sqlboiler/v4/queries/qm"

type GetAlbumArtistsDTO struct {
	AlbumArtist *int
	Limit       *int
	Offset      *int
}

func (d *GetAlbumArtistsDTO) GenWhereMods() []qm.QueryMod {
	mods := make([]qm.QueryMod, 0)

	if d == nil {
		return mods
	}

	if d.AlbumArtist != nil {
		mods = append(mods, qm.Where("album_artist=?", d.AlbumArtist))
	}

	return mods
}
