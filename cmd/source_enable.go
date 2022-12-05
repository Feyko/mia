package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"mia/internal/mia/sources"
)

var sourceEnableCmd = &cobra.Command{
	Use:   "enable",
	Short: "Enable a source for mia to pull from",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")

		err := sources.Enable(name)
		if err != nil {
			log.Fatalf("Could not enable the source: %v", err)
		}
	},
}

func init() {
	sourceCmd.AddCommand(sourceEnableCmd)

	sourceEnableCmd.Flags().StringP("name", "n", "", "Name of the source to enable. Required")
	err := sourceEnableCmd.MarkFlagRequired("name")
	if err != nil {
		log.Fatalf("Couldn't mark the name flag as required: %v", err)
	}
}
