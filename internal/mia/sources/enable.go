package sources

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
)

func Enable(src string) error {
	if src == "all" {
		for _, source := range SupportedSources {
			Enable(source)
		}
		return nil
	}

	found, _ := sourceIn(src, SupportedSources)
	if !found {
		return errors.New("This source is not supported. See 'mia source list'")
	}

	enabledSrcs := viper.GetStringSlice("enabledSources")
	fmt.Printf("source enable: current config is %v\n", enabledSrcs)
	found, _ = sourceIn(src, enabledSrcs)
	if found {
		return errors.New("This source is already enabled")
	}

	enabledSrcs = append(enabledSrcs, src)
	fmt.Printf("source enable: config after change is %v\n", enabledSrcs)
	viper.Set("enabledSources", enabledSrcs)
	err := viper.WriteConfig()
	if err != nil {
		return err
	}

	return nil
}
