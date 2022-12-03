package sources

import (
	"errors"
)

var SupportedSources = []string{
	"nyaan",
	"piratebay",
}

func FormattedList() (output string) {
	output += "FormattedList of the sources currently supported:\n"
	for _, source := range SupportedSources {
		output += ("-" + source + "\n")
	}
	return
}

func sourceIn(src string, list []string) (bool, int) {
	for i, srcInList := range list {
		if src == srcInList {
			return true, i
		}
	}
	return false, -1
}

func removeSource(sources []string, src string) ([]string, error) {
	exists, i := sourceIn(src, sources)
	if !exists {
		return sources, errors.New("The source could not be found in the slice")
	}
	sources[i] = sources[len(sources)-1]
	return sources[:len(sources)-1], nil
}