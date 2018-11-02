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

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "register a user with username ,password , email and telephone number",
	Long: `Usage example: 
	register -uUsername -password pass -email=a@xxx.com -phone=xxxxxxxx
	 `,
	Run: func(cmd *cobra.Command, args []string) {
		tmp_u, _ := cmd.Flags().GetString("username")
		tmp_p, _ := cmd.Flags().GetString("password")
		tmp_m, _ := cmd.Flags().GetString("email")
		tmp_c, _ := cmd.Flags().GetString("phone")
		if tmp_u == "" || tmp_p == "" || tmp_m == "" || tmp_c == "" {
			fmt.Println("Please tell us your username[-u], password[-p], email[-e], cellphone[-c]")
			return
		}
		pass, err := entity.UserRegister(tmp_u, tmp_p, tmp_m, tmp_c)
		if pass == false {
			fmt.Println("Username existed!")
			return
		} else {
			if err != nil {
				fmt.Println("Some unexpected error happened when try to record your info,Please read error.log for detail")
				return
			} else {
				fmt.Println("Successfully register!")
			}
		}
		fmt.Println("register called")
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)

	registerCmd.Flags().StringP("username", "u", "Anonymous", "username")
	registerCmd.Flags().StringP("password", "p", "", "password")
	registerCmd.Flags().StringP("email", "e", "", "email like xxx@xx.com")
	registerCmd.Flags().StringP("phone", "t", "", "telephone number")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// registerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// registerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
