package repository

import (
	"context"
	"database/sql"
	"github.com/legato/domain/errors"
	"github.com/legato/infrastructure/database/dbmodel"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type TrackRepository interface {
	GetByID(ctx context.Context, trackID int64) (*dbmodel.Track, error)
	GetByFilePath(ctx context.Context, filePath string) (*dbmodel.Track, error)
}

func NewTrackRepository(db sqlExecutor) TrackRepository {
	return &trackRepository{
		db: db,
	}
}

type trackRepository struct {
	db sqlExecutor
}

func (r *trackRepository) GetByID(ctx context.Context, trackID int64) (*dbmodel.Track, error) {
	track, err := dbmodel.FindTrack(ctx, r.db, int(trackID))
	if err != nil {
		if err == sql.ErrNoRows {
			errors.New(errors.TrackNotFoundError, nil)
		}
		return nil, err
	}
	return track, nil
}

func (r *trackRepository) GetByFilePath(ctx context.Context, filePath string) (*dbmodel.Track, error) {
	tracks, err := dbmodel.Tracks(qm.Where("file_path_hash=?", genHash(filePath))).All(ctx, r.db)
	if err != nil {
		return nil, err
	}
	if len(tracks) == 0 {
		return nil, errors.New(errors.TrackNotFoundError, nil)
	}
	return tracks[0], err
}