package sources

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
)

func Disable(src string) error {
	if src == "all" {
		for _, source := range SupportedSources {
			Disable(source)
		}
		return nil
	}


	found, _ := sourceIn(src, SupportedSources)
	if !found {
		return errors.New("This source is not supported. See 'mia source list'")
	}

	enabledSrcs := viper.GetStringSlice("enabledSources")
	fmt.Printf("source disable: current config is %v\n", enabledSrcs)
	enabledSrcs, err := removeSource(enabledSrcs, src)
	if err != nil {
		return errors.New("This source is already disabled")
	}

	fmt.Printf("source disable: config after change is %v\n", enabledSrcs)
	viper.Set("enabledSources", enabledSrcs)
	err = viper.WriteConfig()
	if err != nil {
		return err
	}

	return nil
}