/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"mia/internal/mia/sources"
)

// sourceDisableCmd represents the sourceEnable command
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
