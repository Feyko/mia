package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"mia/internal/mia/sources"
)

var sourceListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists the supported sources",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(sources.FormattedList())
	},
}

func init() {
	sourceCmd.AddCommand(sourceListCmd)
}
