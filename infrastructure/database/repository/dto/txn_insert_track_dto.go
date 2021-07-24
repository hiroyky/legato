package dto

import "github.com/legato/infrastructure/database/dbmodel"

type TxnInsertTrackDTO struct {
	Track       *dbmodel.Track
	Album       *dbmodel.Album
	AlbumArtist *dbmodel.AlbumArtist
	Genre       *dbmodel.Genre
}
