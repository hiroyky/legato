package repository

import (
	"context"
	"database/sql"
	"github.com/hiroyky/legato/domain/errors"
	"github.com/hiroyky/legato/infrastructure/database/dbmodel"
	"github.com/hiroyky/legato/infrastructure/database/repository/dto"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type AlbumArtistRepository interface {
	GetByID(ctx context.Context, albumArtistID int64) (*dbmodel.AlbumArtist, error)
	GetByName(ctx context.Context, name string) (*dbmodel.AlbumArtist, error)
	GetAlbumArtists(ctx context.Context, data *dto.GetAlbumArtistsDTO) (dbmodel.AlbumArtistSlice, error)
	CountAlbumArtists(ctx context.Context, data *dto.GetAlbumArtistsDTO) (int64, error)
}

func NewAlbumArtistRepository(db sqlExecutor) AlbumArtistRepository {
	return &albumArtistRepository{db: db}
}

type albumArtistRepository struct {
	db sqlExecutor
}

func (r *albumArtistRepository) GetByID(ctx context.Context, albumArtistID int64) (*dbmodel.AlbumArtist, error) {
	albumArtist, err := dbmodel.FindAlbumArtist(ctx, r.db, int(albumArtistID))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New(errors.AlbumArtistNotFoundError, err)
		}
		return nil, err
	}
	return albumArtist, nil
}

func (r *albumArtistRepository) GetByName(ctx context.Context, name string) (*dbmodel.AlbumArtist, error) {
	albumArtists, err := dbmodel.AlbumArtists(qm.Where("name_hash=?", genHash(name))).All(ctx, r.db)
	if err != nil {
		return nil, err
	}
	if len(albumArtists) == 0 {
		return nil, errors.New(errors.AlbumArtistNotFoundError, nil)
	}
	return albumArtists[0], nil
}

func (r *albumArtistRepository) GetAlbumArtists(ctx context.Context, data *dto.GetAlbumArtistsDTO) (dbmodel.AlbumArtistSlice, error) {
	albumArtists, err := dbmodel.AlbumArtists(appendLimitOffsetMods(data.GenWhereMods(), data.Limit, data.Offset)...).All(ctx, r.db)
	if err != nil {
		return nil, errors.New(errors.GetAlbumArtistFatal, err)
	}
	return albumArtists, nil
}

func (r *albumArtistRepository) CountAlbumArtists(ctx context.Context, data *dto.GetAlbumArtistsDTO) (int64, error) {
	total, err := dbmodel.AlbumArtists(data.GenWhereMods()...).Count(ctx, r.db)
	if err != nil {
		return total, errors.New(errors.CountAlbumArtistFatal, err)
	}
	return total, nil
}
