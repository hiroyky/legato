package registry

import "github.com/legato/infrastructure/database/repository"

func NewTrackRepository() repository.TrackRepository {
	return repository.NewTrackRepository(LegatoDB)
}

func NewAlbumRepository() repository.AlbumRepository {
	return repository.NewAlbumRepository(LegatoDB)
}

func NewAlbumArtistRepository() repository.AlbumArtistRepository {
	return repository.NewAlbumArtistRepository(LegatoDB)
}

func NewGenreRepository() repository.GenreRepository {
	return repository.NewGenreRepository(LegatoDB)
}

func NewTxnInsertTrack() repository.TxnInsertTrack {
	return repository.NewTxnInsertTrack(LegatoDB)
}
