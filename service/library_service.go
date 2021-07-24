package service

import (
	"context"
	"github.com/dhowden/tag"
	"github.com/ktnyt/go-moji"
	"github.com/legato/infrastructure/database/dbmodel"
	"github.com/legato/infrastructure/database/repository"
	"github.com/legato/infrastructure/database/repository/dto"
)

type LibraryService interface {
	InsertTrack(ctx context.Context, metadata tag.Metadata, fileHash, filePath string) error
}

func NewLibraryService(txnInsertTrack repository.TxnInsertTrack) LibraryService {
	return &libraryService{
		txnInsertTrack: txnInsertTrack,
	}
}

type libraryService struct {
	txnInsertTrack repository.TxnInsertTrack
}

func (s *libraryService) InsertTrack(ctx context.Context, metadata tag.Metadata, fileHash, filePath string) error {
	trackNo, _ := metadata.Track()
	track := &dbmodel.Track{
		Title:    s.convertMoji(metadata.Title()),
		Artist:   s.convertMoji(metadata.Artist()),
		Composer: s.convertMoji(metadata.Composer()),
		TrackNo:  trackNo,
		Lyrics:   metadata.Lyrics(),
		Comment:  metadata.Comment(),
		Year:     metadata.Year(),
		FilePath: filePath,
		FileHash: fileHash,
		Format:   string(metadata.Format()),
		FileType: string(metadata.FileType()),
	}

	discNo, discTotal := metadata.Disc()
	album := &dbmodel.Album{
		Name:      s.convertMoji(metadata.Album()),
		DiscNo:    discNo,
		DiscTotal: discTotal,
	}
	albumArtist := &dbmodel.AlbumArtist{
		Name: s.convertMoji(metadata.AlbumArtist()),
	}
	genre := &dbmodel.Genre{
		Name: s.convertMoji(metadata.Genre()),
	}

	return s.txnInsertTrack.Commit(ctx, &dto.TxnInsertTrackDTO{
		Track:       track,
		Album:       album,
		AlbumArtist: albumArtist,
		Genre:       genre,
	})
}

func (s *libraryService) convertMoji(src string) string {
	// 全角英数を半角英数に
	dst := moji.Convert(src, moji.ZE, moji.HE)
	// 半角カタカナを全角カタカナに
	dst = moji.Convert(dst, moji.HK, moji.ZK)
	// 全角スペースを半角スペースに
	dst = moji.Convert(dst, moji.ZS, moji.ZS)
	return dst
}
