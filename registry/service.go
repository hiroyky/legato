package registry

import (
	"context"
	"github.com/hiroyky/legato/service"
)

func NewMetadataService(_ context.Context) service.MetadataService {
	return service.NewMetadataService()
}

func NewLibraryService(_ context.Context) service.LibraryService {
	return service.NewLibraryService(NewTxnInsertTrack())
}
