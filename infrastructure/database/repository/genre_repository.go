package repository

import (
	"context"
	"github.com/legato/domain/errors"
	"github.com/legato/infrastructure/database/dbmodel"
	"github.com/legato/infrastructure/database/repository/dto"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type GenreRepository interface {
	GetByID(ctx context.Context, genreID int64) (*dbmodel.Genre, error)
	GetByName(ctx context.Context, name string) (*dbmodel.Genre, error)
	GetGenres(ctx context.Context, data *dto.GetGenresDTO) (dbmodel.GenreSlice, error)
	CountGenres(ctx context.Context, data *dto.GetGenresDTO) (int64, error)
}

func NewGenreRepository(db sqlExecutor) GenreRepository {
	return &genreRepository{db: db}
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

func (r *genreRepository) GetGenres(ctx context.Context, data *dto.GetGenresDTO) (dbmodel.GenreSlice, error) {
	genres, err := dbmodel.Genres(appendLimitOffsetMods(data.GenWhereMods(), data.Limit, data.Offset)...).All(ctx, r.db)
	if err != nil {
		return nil, errors.New(errors.GetGenreFatal, err)
	}
	return genres, nil
}

func (r *genreRepository) CountGenres(ctx context.Context, data *dto.GetGenresDTO) (int64, error) {
	total, err := dbmodel.Genres(data.GenWhereMods()...).Count(ctx, r.db)
	if err != nil {
		return total, errors.New(errors.CountGenreFatal, err)
	}
	return total, nil
}
