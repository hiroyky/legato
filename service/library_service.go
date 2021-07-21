package service

import "github.com/dhowden/tag"

type LibraryService interface {
	InsertTrack(metadata tag.Metadata, fileMD5, filePath string) error
}

func NewLibraryService() LibraryService {
	return &libraryService{}
}

type libraryService struct {
}

func (s *libraryService) InsertTrack(metadata tag.Metadata, fileMD5, filePath string) error {
	return nil
}
