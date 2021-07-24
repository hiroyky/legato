package registry

import (
	"context"
	"github.com/legato/service"
)

func NewMetadataService(_ context.Context) service.MetadataService {
	return service.NewMetadataService()
}

func NewLibraryService(_ context.Context) service.LibraryService {
	return service.NewLibraryService(NewTxnInsertTrack())
}
