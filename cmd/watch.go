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
	"mia/internal/mia"
)

// watchCmd represents the watch command
var watchCmd = &cobra.Command{
	Use:   "watch",
	Short: "Schedule mia to be ran at an interval. Default is every day",
	Run: run,
}

func init() {
	rootCmd.AddCommand(watchCmd)

	watchCmd.Flags().IntP("interval", "i",1, "The interval in days at which mia will run")
}

func run(cmd *cobra.Command, args []string) {
	interval, _ := cmd.Flags().GetInt("interval")
	mia.Watch(interval)
}