package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var sourceCmd = &cobra.Command{
	Use:   "source",
	Short: "Manage the sources mia pulls from",
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		if err != nil {
			log.Fatalf("Could not display the help message: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(sourceCmd)
}
