package service

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/dhowden/tag"
	"github.com/friendsofgo/errors"
	"io"
	"os"
)

type MetadataService interface {
	ParseMetadata(fileName string) (tag.Metadata, error)
	ParseFileMD5(fileName string) (string, error)
}

func NewMetadataService() MetadataService {
	return &metadataService{}
}

type metadataService struct {
}

func (s *metadataService) ParseMetadata(fileName string) (tag.Metadata, error) {
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		return nil, err
	}

	metadata, err := tag.ReadFrom(file)
	if err != nil {
		if err == tag.ErrNoTagsFound {
			return nil, nil
		}
		return nil, err
	}

	return metadata, nil
}

func (s *metadataService) ParseFileMD5(fileName string) (string, error) {
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		return "", err
	}

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", errors.Wrap(err, fileName)
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}
