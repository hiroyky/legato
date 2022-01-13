package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/friendsofgo/errors"
	"github.com/hiroyky/legato/lib"
	"github.com/hiroyky/legato/registry"
	"io/ioutil"
	"path/filepath"
	"strings"
)

var metadataService = registry.NewMetadataService(context.Background())
var libraryService = registry.NewLibraryService(context.Background())

func main() {
	baseDir := flag.String("base-dir", "/mnt/music", "root directory for search.")
	ext := flag.String("ext", ".mp3,.flac,.aac", "extensions of target file. multiple specification is available by comma split.")
	flag.Parse()

	extensions := strings.Split(*ext, ",")

	fmt.Printf("Search music files from %s \n", *baseDir)
	fmt.Printf("Extensions of target files are %#v", extensions)

	if err := insert(context.Background(), *baseDir, extensions); err != nil {
		panic(err)
	}
}

func insert(ctx context.Context, baseDir string, extensions []string) error {
	fmt.Printf("Searching at %s \n", baseDir)
	contents, err := ioutil.ReadDir(baseDir)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Fatal to readDir at %s", baseDir))
	}

	for _, c := range contents {
		p := filepath.Join(baseDir, c.Name())
		if c.IsDir() {
			if err := insert(ctx, p, extensions); err != nil {
				return err
			}
			continue
		}

		if !lib.ContainStr(filepath.Ext(c.Name()), extensions) {
			continue
		}

		fmt.Printf("Import %s \n", p)
		meta, err := metadataService.ParseMetadata(p)
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("Fatal to parse metadata %s", p))
		}
		md5Hash, err := metadataService.ParseFileMD5(p)
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("Fatal to parse md5 hash %s", p))
		}

		if meta == nil {
			if err := libraryService.InsertTrackWithoutMetadata(ctx, md5Hash, p); err != nil {
				return errors.Wrap(err, fmt.Sprintf("Fatal to insert track %s", p))
			}
		} else {
			if err := libraryService.InsertTrack(ctx, meta, md5Hash, p); err != nil {
				return errors.Wrap(err, fmt.Sprintf("Fatal to insert track %s", p))
			}
		}
	}
	return nil
}
