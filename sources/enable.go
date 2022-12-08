package sources

import (
	"github.com/pkg/errors"
)

func Enable(src string) error {
	if src == "all" {
		for _, source := range SupportedSources {
			err := Enable(source)
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

	return enableSource(src)
}
