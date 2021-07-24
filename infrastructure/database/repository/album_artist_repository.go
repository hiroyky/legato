package repository

import (
	"context"
	"database/sql"
	"github.com/legato/domain/errors"
	"github.com/legato/infrastructure/database/dbmodel"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type AlbumArtistRepository interface {
	GetByID(ctx context.Context, albumArtistID int64) (*dbmodel.AlbumArtist, error)
	GetByName(ctx context.Context, name string) (*dbmodel.AlbumArtist, error)
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
