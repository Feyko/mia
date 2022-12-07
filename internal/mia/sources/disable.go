package sources

import (
	"github.com/pkg/errors"
)

func Disable(src string) error {
	if src == "all" {
		for _, source := range SupportedSources {
			err := Disable(source)
			if err != nil {
				return err
			}
		}
		return nil
	}

	supported := IsSupported(src)
	if !supported {
		return errors.Errorf("The source %q is not supported. See 'mia source list'", src)
	}

	return disableSource(src)
}
