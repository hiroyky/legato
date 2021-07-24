package repository

import (
	"context"
	"database/sql"
	"github.com/friendsofgo/errors"
	"github.com/legato/infrastructure/database/dbmodel"
	"github.com/legato/infrastructure/database/repository/dto"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type TxnInsertTrack interface {
	Commit(ctx context.Context, o *dto.TxnInsertTrackDTO) error
}

func NewTxnInsertTrack(db sqlExecutor) TxnInsertTrack {
	return &txnInsertTrack{db: db}
}

type txnInsertTrack struct {
	db sqlExecutor
}

func (t *txnInsertTrack) Commit(ctx context.Context, o *dto.TxnInsertTrackDTO) error {
	tx, err := t.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	if err := t.transaction(ctx, tx, o); err != nil {
		if err := tx.Rollback(); err != nil {
			return errors.Wrap(err, "Failed to rollback")
		}
		return errors.Wrap(err, "Rollback: "+o.Track.FilePath)
	}

	return tx.Commit()
}

func (t *txnInsertTrack) transaction(ctx context.Context, tx *sql.Tx, o *dto.TxnInsertTrackDTO) error {
	o.Genre.NameHash = genHash(o.Genre.Name)
	existGenre, err := dbmodel.Genres(qm.Where("name_hash=?", o.Genre.NameHash)).One(ctx, tx)
	if err == sql.ErrNoRows {
		if err := o.Genre.Insert(ctx, tx, boil.Infer()); err != nil {
			return err
		}
	} else if err != nil {
		return err
	}
	if existGenre != nil {
		o.Genre.GenreID = existGenre.GenreID
		if _, err := o.Genre.Update(ctx, tx, boil.Infer()); err != nil {
			return err
		}
	}
	o.Track.GenreID = o.Genre.GenreID
	o.AlbumArtist.NameHash = genHash(o.AlbumArtist.Name)
	existAlbumArtist, err := dbmodel.AlbumArtists(qm.Where("name_hash=?", o.AlbumArtist.NameHash)).One(ctx, tx)
	if err == sql.ErrNoRows {
		if err := o.AlbumArtist.Insert(ctx, tx, boil.Infer()); err != nil {
			return err
		}
	} else if err != nil {
		return err
	}
	if existAlbumArtist != nil {
		o.AlbumArtist.AlbumArtistID = existAlbumArtist.AlbumArtistID
		if _, err := o.AlbumArtist.Update(ctx, tx, boil.Infer()); err != nil {
			return err
		}
	}
	o.Track.AlbumArtistID = o.AlbumArtist.AlbumArtistID
	o.Album.AlbumArtistID = o.AlbumArtist.AlbumArtistID
	existAlbum, err := dbmodel.Albums(
		qm.Where("name=?", o.Album.Name),
		qm.Where("disc_no=?", o.Album.DiscNo),
		qm.WhereIn("disc_total=?", o.Album.DiscTotal),
	).One(ctx, tx)
	if err == sql.ErrNoRows {
		if err := o.Album.Insert(ctx, tx, boil.Infer()); err != nil {
			return err
		}
	} else if err != nil {
		return err
	}
	if existAlbum != nil {
		o.Album.AlbumID = existAlbum.AlbumID
		if _, err := o.Album.Update(ctx, tx, boil.Infer()); err != nil {
			return err
		}
	}
	o.Track.AlbumID = o.Album.AlbumID

	o.Track.FilePathHash = genHash(o.Track.FilePath)
	existTrack, err := dbmodel.Tracks(qm.Where("file_path_hash=?", o.Track.FilePathHash)).One(ctx, tx)
	if err == sql.ErrNoRows {
		if err := o.Track.Insert(ctx, tx, boil.Infer()); err != nil {
			return err
		}
	} else if err != nil {
		return err
	}
	if existTrack != nil {
		o.Track.TrackID = existTrack.TrackID
		if _, err := o.Track.Update(ctx, tx, boil.Infer()); err != nil {
			return err
		}
	}
	return nil
}
