package sources

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"golang.org/x/exp/slices"
)

var SupportedSources = []string{
	"nyaa",
	"piratebay",
}

func FormattedList() (output string) {
	output += "List of the sources currently supported:\n"
	for _, source := range SupportedSources {
		output += ("-" + source + "\n")
	}
	return
}

func IsSupported(src string) bool {
	return slices.Contains(SupportedSources, src)
}

func IsEnabled(src string) bool {
	enabledSrcs := viper.GetStringSlice("enabledSources")
	return slices.Contains(enabledSrcs, src)
}

func enableSource(src string) error {
	if !IsSupported(src) {
		return errors.Errorf("unsupported source %q", src)
	}
	if IsEnabled(src) {
		return nil
	}
	enabledSrcs := viper.GetStringSlice("enabledSources")
	enabledSrcs = append(enabledSrcs, src)
	viper.Set("enabledSources", enabledSrcs)
	err := viper.WriteConfig()
	if err != nil {
		return errors.Wrap(err, "error writing config to disk")
	}
	return nil
}

func disableSource(src string) error {
	if !IsSupported(src) {
		return errors.Errorf("unsupported source %q", src)
	}
	if !IsEnabled(src) {
		return nil
	}

	enabledSrcs := viper.GetStringSlice("enabledSources")
	enabledSrcs, err := removeSource(enabledSrcs, src)
	if err != nil {
		return errors.Wrap(err, "error removing source")
	}

	viper.Set("enabledSources", enabledSrcs)
	err = viper.WriteConfig()
	if err != nil {
		return errors.Wrap(err, "error writing config to disk")
	}
	return nil
}

func removeSource(sources []string, src string) ([]string, error) {
	i := slices.Index(sources, src)
	if i < 0 {
		return sources, nil
	}
	sources[i] = sources[len(sources)-1]
	return sources[:len(sources)-1], nil
}
