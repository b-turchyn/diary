/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

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
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/b-turchyn/diary/lib"
	"github.com/spf13/cobra"
)

type LogType struct {
	DbName string
	Header string
}

// todayCmd represents the today command
var todayCmd = &cobra.Command{
	Use:   "today",
	Short: "Today's entries",
	Long: `Get an output of all entries made on a day. With no parameters, outputs
today's entries.

You can specify a specific date using --date in ISO8601 date format.`,
	Run: func(cmd *cobra.Command, args []string) {
		logs := []LogType{
			{DbName: "mistakes", Header: "On This Day I Screwed Up"},
			{DbName: "log", Header: "On This Day I Did"},
			{DbName: "learn", Header: "On This Day I Learned"},
		}
		db := lib.NewDB()

		var date time.Time
		if overrideDate != "" {
			var err error
			date, err = time.Parse("2006-01-02", overrideDate)

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		} else {
			date = time.Now()
		}
		// Truncate the time and ensure we are working in the local timezone
		date = time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Local().Location())

		output := []string{
			fmt.Sprintf(lib.PrimaryText("# Notes For %s\n"), date.Format("Monday Jan 2, 2006")),
		}

		for _, v := range logs {
			log, err := lib.GetLogBlock(db, v.DbName, v.Header, date)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			o := log.ToString()
			if o != "" {
				output = append(output, o)
			}
		}

		fmt.Println(strings.Join(output, "\n"))
	},
}

func init() {
	rootCmd.AddCommand(todayCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// todayCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// todayCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
