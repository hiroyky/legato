package repository

import (
	"context"
	"database/sql"
	"github.com/hiroyky/legato/domain/errors"
	"github.com/hiroyky/legato/infrastructure/database/dbmodel"
	"github.com/hiroyky/legato/infrastructure/database/repository/dto"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type AlbumRepository interface {
	GetByID(ctx context.Context, albumID int) (*dbmodel.Album, error)
	GetByName(ctx context.Context, name string) (*dbmodel.Album, error)
	GetAlbums(ctx context.Context, data *dto.GetAlbumsDTO) (dbmodel.AlbumSlice, error)
	CountAlbums(ctx context.Context, data *dto.GetAlbumsDTO) (int64, error)
}

func NewAlbumRepository(db sqlExecutor) AlbumRepository {
	return &albumRepository{db: db}
}

type albumRepository struct {
	db sqlExecutor
}

func (r *albumRepository) GetByID(ctx context.Context, albumID int) (*dbmodel.Album, error) {
	album, err := dbmodel.FindAlbum(ctx, r.db, int(albumID))
	if err != nil {
		return nil, errors.New(errors.GetAlbumFatal, err)
	}
	return album, nil
}

func (r *albumRepository) GetByName(ctx context.Context, name string) (*dbmodel.Album, error) {
	album, err := dbmodel.Albums(qm.Where("name_hash=?", genHash(name))).One(ctx, r.db)
	if err == sql.ErrNoRows {
		return nil, errors.New(errors.AlbumNotFoundError, nil)
	}
	if err != nil {
		return nil, err
	}
	return album, nil
}

func (r *albumRepository) GetAlbums(ctx context.Context, data *dto.GetAlbumsDTO) (dbmodel.AlbumSlice, error) {
	res, err := dbmodel.Albums(appendLimitOffsetMods(data.GenWhereMods(), data.Limit, data.Offset)...).All(ctx, r.db)
	if err != nil {
		return nil, errors.New(errors.GetAlbumFatal, err)
	}
	return res, nil
}

func (r *albumRepository) CountAlbums(ctx context.Context, data *dto.GetAlbumsDTO) (int64, error) {
	res, err := dbmodel.Albums(data.GenWhereMods()...).Count(ctx, r.db)
	if err != nil {
		return 0, errors.New(errors.CountAlbumFatal, err)
	}
	return res, nil
}
