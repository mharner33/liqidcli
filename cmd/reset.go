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
	"bytes"
	"fmt"
	"net/http"
	"strings"

	"github.com/mharner33/liqidcli/lapi/liqutil"
	"github.com/spf13/cobra"
)

// resetCmd represents the reset command
var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Perform a reset on the Liqid environment",
	Long: `Removes all device configuration information (groups, machines etc.) and 
	returns them back to the free pool. For example:

	liqidcli --ip 10.204.103.38 reset`,
	Run: func(cmd *cobra.Command, args []string) {
		var ack string
		fmt.Print("Are you sure you want to perform a reset (y/n)?")
		fmt.Scanln(&ack)
		if strings.ToLower(ack) == "y" {
			fmt.Println("Resetting the system...")
			_, err := http.Post("http://"+ipAddress+liqutil.ApiPath+"system/state/reset", "application/json", bytes.NewBuffer(nil))
			if err != nil {
				fmt.Printf("Post request for system reset failed: %s\n", err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(resetCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// resetCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// resetCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
