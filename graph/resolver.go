package graph

import "github.com/hiroyky/legato/infrastructure/database/repository"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	TrackRepository       repository.TrackRepository
	AlbumRepository       repository.AlbumRepository
	AlbumArtistRepository repository.AlbumArtistRepository
	GenreRepository       repository.GenreRepository
}
