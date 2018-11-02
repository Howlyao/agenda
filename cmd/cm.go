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

// cmCmd represents the cm command
var cmCmd = &cobra.Command{
	Use:   "createmeeting",
	Short: "create meeting command",
	Run: func(cmd *cobra.Command, args []string) {

		title, _ := cmd.Flags().GetString("title")
		participator, _ := cmd.Flags().GetStringSlice("participator")
		starttime, _ := cmd.Flags().GetString("starttime")
		endtime, _ := cmd.Flags().GetString("endtimeendtime")
		if title == "" || len(participator) == 0 || starttime == "" || endtime == "" {
			fmt.Println("Please input title, starttime[yyyy-mm-dd/hh:mm],endtime,participator(input like \"name1, name2\")")
			return
		}
		if user, flag := entity.GetCurUser(); flag != true {
			fmt.Println("Error: please login firstly")
			return
		} else {
			// participators := strings.Split(tmp_p,",")
			if flag := entity.CreateMeeting(user.Username, title, starttime, endtime, participator); flag != true {
				fmt.Println("Error: create Failed. Please check error.log for more detail")
				return
			} else {
				fmt.Println("Create meeting successfully!")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(cmCmd)

	// Here you will define your flags and configuration settings.
	cmCmd.Flags().StringP("title", "t", "", "the title of meeting")
	cmCmd.Flags().StringSliceP("participator", "p", nil, "the participator(s) of the meeting, split by comma,input like \"name1, name2\"")
	cmCmd.Flags().StringP("starttime", "s", "", "the startTime of the meeting")
	cmCmd.Flags().StringP("endtime", "e", "", "the endTime of the meeting")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
