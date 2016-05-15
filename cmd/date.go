// Copyright 息 2016 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

// dateCmd represents the date command
var dateCmd = &cobra.Command{
	Use:   "show-date",
	Short: "show date and time simply",
	Long: `show-date is just simple displaying date and time.
You can specify the number to add or subtract seconds as argument.`,
	Run: func(cmd *cobra.Command, args []string) {
		PrintDate(AddFlag, SubFlag)
	},
}

var AddFlag, SubFlag int

func PrintDate(add, sub int) {
	now := time.Now()
	added := now.Add(time.Duration(add) * time.Second)
	result := added.Add(time.Duration(-sub) * time.Second)

	fmt.Println(result.Format(time.RFC3339))
}

func init() {
	RootCmd.AddCommand(dateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dateCmd.PersistentFlags().String("foo", "", "A help for foo")
	dateCmd.PersistentFlags().IntVarP(&AddFlag, "add", "a", 0, "Add the seconds to time")
	dateCmd.PersistentFlags().IntVarP(&SubFlag, "sub", "s", 0, "Add the seconds to time")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
