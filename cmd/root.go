package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"mia/internal/mia"
	"mia/internal/mia/sources"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "mia",
	Short: "Easily download and watch releases of a multitude of media content",
	Long: `Mia is a tool that lets you easily download and watch releases of a multitude of media content
Once configured, simply use 'mia'
You can also use 'mia watch' to ask mia to mia periodically`,
	Run: func(cmd *cobra.Command, args []string) { mia.Run() },
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.mia.yaml)")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.SetConfigName(".mia")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(home)
	}

	viper.SetDefault("enabledSources", sources.SupportedSources)

	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err == nil {
		return
	}

	err = viper.SafeWriteConfig()
	if err != nil {
		log.Fatalf("Couldn't write config file: %v", err)
	}
}
