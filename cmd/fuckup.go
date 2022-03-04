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
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/b-turchyn/diary/lib"
	"github.com/spf13/cobra"
)

// fuckupCmd represents the fuckup command
var fuckupCmd = &cobra.Command{
	Aliases: []string{"fuckup"},
	Use:     "mistake",
	Short:   "You messed up. Explain why.",
	Long: `You did something wrong. Write it out. Get it out of your system so you can learn
	and grow from it.`,
	Args: dateTimeArgsValidator,
	Run: func(cmd *cobra.Command, args []string) {
		t := lib.GetTime(overrideDate, overrideTime)
		input, err := lib.PromptForInput("What'd you do? What'd you learn? ", args)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		db := lib.NewDB()
		err = insertFuckup(db, t, input)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		} else {
			fmt.Println(lib.PrimaryText("Recorded:"), input)
		}
	},
}

func init() {
	rootCmd.AddCommand(fuckupCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fuckupCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fuckupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func insertFuckup(db *sql.DB, t time.Time, text string) error {
	return lib.InsertGenericText(db, "mistakes", t, text)
}
