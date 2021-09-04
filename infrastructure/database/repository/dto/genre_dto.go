package dto

import "github.com/volatiletech/sqlboiler/v4/queries/qm"

type GetGenresDTO struct {
	GenreID *int
	Limit   *int
	Offset  *int
}

func (d *GetGenresDTO) GenWhereMods() []qm.QueryMod {
	mods := make([]qm.QueryMod, 0)

	if d == nil {
		return mods
	}
	if d.GenreID != nil {
		mods = append(mods, qm.Where("genre_id=?", d.GenreID))
	}

	return mods
}
