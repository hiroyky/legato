package repository

import (
	"context"
	"github.com/legato/domain/errors"
	"github.com/legato/infrastructure/database/dbmodel"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type GenreRepository interface {
	GetByID(ctx context.Context, genreID int64) (*dbmodel.Genre, error)
	GetByName(ctx context.Context, name string) (*dbmodel.Genre, error)
}

type genreRepository struct {
	db sqlExecutor
}

func (r *genreRepository) GetByID(ctx context.Context, genreID int64) (*dbmodel.Genre, error) {
	genre, err := dbmodel.FindGenre(ctx, r.db, int(genreID))
	if err != nil {

	}
	return genre, nil
}

func (r *genreRepository) GetByName(ctx context.Context, name string) (*dbmodel.Genre, error) {
	genres, err := dbmodel.Genres(qm.Where("name_hash=?", genHash(name))).All(ctx, r.db)
	if err != nil {
		return nil, err
	}
	if len(genres) == 0 {
		return nil, errors.New(errors.GenreNotFoundError, nil)
	}
	return genres[0], nil
}
