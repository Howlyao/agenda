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

// quCmd represents the qu command
var quCmd = &cobra.Command{
	Use:   "queryuser",
	Short: "To query user",
	Run: func(cmd *cobra.Command, args []string) {

		if _, flag := entity.GetCurUser(); flag != true {
			fmt.Println("Please Log in firstly")
		}
		ru := entity.ListAllUser()
		for _, u := range ru {
			fmt.Println("----------------")
			fmt.Println("Username: ", u.Username)
			fmt.Println("Phone: ", u.Phone)
			fmt.Println("Email: ", u.Email)
			fmt.Println("----------------")
		}
	},
}

func init() {
	rootCmd.AddCommand(quCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// quCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// quCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
