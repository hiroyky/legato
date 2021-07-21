package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/friendsofgo/errors"
	"github.com/legato/lib"
	"github.com/legato/registry"
	"io/ioutil"
	"path/filepath"
	"strings"
)

var metadataService = registry.NewMetadataService(context.Background())
var libraryService = registry.NewLibraryService(context.Background())

func main() {
	baseDir := flag.String("base-dir", "ssss", "root directory for search.")
	ext := flag.String("ext", ".mp3,.flac,aac", "extensions of target file. multiple specification is available by comma split.")
	flag.Parse()

	extensions := strings.Split(*ext, ",")

	fmt.Printf("Search music files from %s \n", *baseDir)
	fmt.Printf("Extensions of target files are %#v", extensions)
}

func insert(baseDir string, extensions []string) error {
	fmt.Printf("Searching at %s", baseDir)
	contents, err := ioutil.ReadDir(baseDir)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Fatal to readDir at %s", baseDir))
	}

	for _, c := range contents {
		p := filepath.Join(baseDir, c.Name())
		if c.IsDir() {
			if err := insert(p, extensions); err != nil {
				return err
			}
			continue
		}

		if !lib.ContainStr(filepath.Ext(c.Name()), extensions) {
			continue
		}

		meta, err := metadataService.ParseMetadata(p)
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("Fatal to parse metadata %s", p))
		}
		md5Hash, err := metadataService.ParseFileMD5(p)
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("Fatal to parse md5 hash %s", p))
		}

		if err := libraryService.InsertTrack(meta, md5Hash, p); err != nil {
			return errors.Wrap(err, fmt.Sprintf("Fatal to insert track %s", p))
		}
	}
	return nil
}
