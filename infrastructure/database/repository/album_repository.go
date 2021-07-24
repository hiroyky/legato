package repository

import (
	"context"
	"github.com/legato/domain/errors"
	"github.com/legato/infrastructure/database/dbmodel"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type AlbumRepository interface {
	GetByID(ctx context.Context, albumID int64) (*dbmodel.Album, error)
	GetByName(ctx context.Context, name string) (*dbmodel.Album, error)
}

type albumRepository struct {
	db sqlExecutor
}

func (r *albumRepository) GetByID(ctx context.Context, albumID int64) (*dbmodel.Album, error) {
	album, err := dbmodel.FindAlbum(ctx, r.db, int(albumID))
	if err != nil {

	}
	return album, nil
}

func (r *albumRepository) GetByName(ctx context.Context, name string) (*dbmodel.Album, error) {
	albums, err := dbmodel.Albums(qm.Where("name_hash=?", genHash(name))).All(ctx, r.db)
	if err != nil {
		return nil, err
	}
	if len(albums) == 0 {
		return nil, errors.New(errors.AlbumNotFoundError, nil)
	}
	return albums[0], nil
}
