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

	"github.com/github-user/agenda/entity"
	"github.com/spf13/cobra"
)

// rmpCmd represents the rmp command
var rmpCmd = &cobra.Command{
	Use:   "removeparticipator",
	Short: "remove participator(s)",
	Run: func(cmd *cobra.Command, args []string) {
		title, _ := cmd.Flags().GetString("title")
		participator, _ := cmd.Flags().GetStringSlice("participator")
		if title == "" || len(participator) == 0 {
			fmt.Println("Please input title and participator(s)(input like \"name1, name2\")")
			return
		}
		if user, flag := entity.GetCurUser(); flag != true {
			fmt.Println("Please login firstly")
		} else {
			// participators := strings.Split(participator, ",")
			flag := entity.RemoveMeetingParticipator(user.Username, title, participator)
			if flag != true {
				fmt.Println("Unexpected error. Check error.log for detail")
			} else {
				fmt.Println("Remove successfully")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(rmpCmd)

	// Here you will define your flags and configuration settings.
	rmpCmd.Flags().StringP("title", "t", "", "the title of the meeting")
	rmpCmd.Flags().StringSliceP("participator", "p", nil, "the participator(s) of the meeting, input like \"name1, name2\"")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rmpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rmpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
