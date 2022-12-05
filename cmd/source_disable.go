package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"mia/internal/mia/sources"
)

var sourceDisableCmd = &cobra.Command{
	Use:   "disable",
	Short: "Disable a source mia pulls from",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")

		err := sources.Disable(name)
		if err != nil {
			log.Fatalf("Could not disable the source: %v", err)
		}
	},
}

func init() {
	sourceCmd.AddCommand(sourceDisableCmd)

	sourceDisableCmd.Flags().StringP("name", "n", "", "Name of the source to disable. Required")
	err := sourceDisableCmd.MarkFlagRequired("name")
	if err != nil {
		log.Fatalf("Couldn't mark the name flag as required: %v", err)
	}
}
