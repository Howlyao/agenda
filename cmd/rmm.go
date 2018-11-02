// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
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

	"github.com/Howlyao/agenda/entity"
	"github.com/spf13/cobra"
)

// rmmCmd represents the rmm command
var rmmCmd = &cobra.Command{
	Use:   "rmm",
	Short: "delete meeting",
	Run: func(cmd *cobra.Command, args []string) {

		title, _ := cmd.Flags().GetString("title")
		if title == "" {
			fmt.Println("Error: Please input meeting title")
			return
		}
		if user, flag := entity.GetCurUser(); flag != true {
			fmt.Println("Error: Please login firstly!")
		} else {
			if c := entity.DeleteMeeting(user.Username, title); c == 0 {
				fmt.Println("Error: Meeting not exist or you're not a Sponsor of it")
			} else {
				fmt.Println("Delete Successfully")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(rmmCmd)

	// Here you will define your flags and configuration settings.
	rmmCmd.Flags().StringP("title", "t", "", "the title of meeting")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rmmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rmmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
