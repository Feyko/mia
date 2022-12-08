package cmd

import (
	"github.com/spf13/cobra"
	"mia"
)

var watchCmd = &cobra.Command{
	Use:   "watch",
	Short: "Schedule mia to be ran at an interval. Default is every day",
	Run:   run,
}

func init() {
	rootCmd.AddCommand(watchCmd)

	watchCmd.Flags().IntP("interval", "i", 1, "The interval in days at which mia will run")
}

func run(cmd *cobra.Command, args []string) {
	interval, _ := cmd.Flags().GetInt("interval")
	main.Watch(interval)
}
