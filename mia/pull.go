package mia

import (
	"fmt"
	"github.com/pkg/errors"
	"google.golang.org/appengine/log"
	"io"
	"mia/media"
	"mia/sources"
)

func Pull() error {
	allMedia, err := media.List()
	if err != nil {
		return errors.Wrap(err, "error getting media list")
	}

	sources_ := sources.List()

	for _, m := range allMedia {
		for _, source := range sources_ {
			found, err := source.Search(m.Match)
			if err != nil {
				log.Errorf(nil, "error searching for media %q on source %q", m.Name, source.Name())
			}
			if found {
				reader, err := source.Download(m.Match)
				if err != nil {
					return errors.Wrap(err, "error downloading media")
				}
				defer reader.Close()

				b, err := io.ReadAll(reader)
				if err != nil {
					return errors.Wrap(err, "error reading downloaded file")
				}

				fmt.Printf("Downloaded file of length %v", len(b))
			}
		}
	}
	
	return nil
}
