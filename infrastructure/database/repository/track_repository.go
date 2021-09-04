package repository

import (
	"context"
	"database/sql"
	"github.com/legato/domain/errors"
	"github.com/legato/infrastructure/database/dbmodel"
	"github.com/legato/infrastructure/database/repository/dto"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type TrackRepository interface {
	GetByID(ctx context.Context, trackID int64) (*dbmodel.Track, error)
	GetByFilePath(ctx context.Context, filePath string) (*dbmodel.Track, error)
	GetByFilePathHash(ctx context.Context, filePathHash string) (*dbmodel.Track, error)
	GetTracks(ctx context.Context, data *dto.GetTracksDTO) (dbmodel.TrackSlice, error)
	CountTracks(ctx context.Context, data *dto.GetTracksDTO) (int64, error)
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
	if err == sql.ErrNoRows {
		return nil, errors.New(errors.TrackNotFoundError, nil)
	}
	if err != nil {
		return nil, errors.New(errors.GetTrackFatal, err)
	}
	return track, nil
}

func (r *trackRepository) GetByFilePath(ctx context.Context, filePath string) (*dbmodel.Track, error) {
	track, err := dbmodel.Tracks(qm.Where("file_path_hash=?", genHash(filePath))).One(ctx, r.db)
	if err == sql.ErrNoRows {
		return nil, errors.New(errors.TrackNotFoundError, nil)
	}
	if err != nil {
		return nil, errors.New(errors.GetTrackFatal, err)
	}
	return track, err
}

func (r *trackRepository) GetByFilePathHash(ctx context.Context, filePathHash string) (*dbmodel.Track, error) {
	track, err := dbmodel.Tracks(qm.Where("file_path_hash=?", filePathHash)).One(ctx, r.db)
	if err == sql.ErrNoRows {
		return nil, errors.New(errors.TrackNotFoundError, nil)
	}
	if err != nil {
		return nil, errors.New(errors.GetTrackFatal, err)
	}
	return track, err
}

func (r *trackRepository) GetTracks(ctx context.Context, data *dto.GetTracksDTO) (dbmodel.TrackSlice, error) {
	res, err := dbmodel.Tracks(appendLimitOffsetMods(data.GenWhereMods(), data.Limit, data.Offset)...).All(ctx, r.db)
	if err != nil {
		return nil, errors.New(errors.GetTrackFatal, err)
	}
	return res, nil
}

func (r *trackRepository) CountTracks(ctx context.Context, data *dto.GetTracksDTO) (int64, error) {
	res, err := dbmodel.Tracks(data.GenWhereMods()...).Count(ctx, r.db)
	if err != nil {
		return 0, errors.New(errors.CountTrackFatal, err)
	}
	return res, nil
}
