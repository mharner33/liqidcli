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
	"fmt"
	"strings"

	"github.com/mharner33/liqidcli/lapi/liqtopo"
	"github.com/mharner33/liqidcli/lapi/liqutil"
	"github.com/spf13/cobra"
)

var (
	listType string
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List information about current infrastructure",
	Long: `Gives information about the current Liqid infrastructure including
	information on groups, machines and acclerator resources.  
	
	Valid flags for type:  all, group, machine and version
	
	For example:
	liqidcli --ip 10.204.105.38 list --type all
	liqidcli --ip 10.204.105.38 list --type version`,
	//Example:  liqidcli --ip 10.204.105.38 list --type version

	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Printf("list called with IP: %s", ipAddress)
		basePath := "http://" + ipAddress + liqutil.ApiPath
		fabid := liqtopo.GetFabID(basePath + "fabric/id")
		listType = strings.ToLower(listType)

		switch listType {
		case "version":
			liqtopo.GetVersion(basePath + "version")
		case "group":
			//Get the fabric ID so we can list the groups

			//fmt.Printf("FabricID: %s\n", fabid)
			qstring := "group?parameters=grp_id%3D" + fabid
			liqtopo.GetGroups(basePath, qstring)
		case "resource":
			qstring := "devices/count?fabr_id=" + fabid
			liqtopo.GetDevicCnt(basePath + qstring)
		default:
			fmt.Printf("Bad argument supplied to list command: %s\n", listType)
		}
		//liqtopo.ListSwitch(basePath, listType)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().StringVarP(&listType, "type", "t", "", "The type should be one of: all, group, machine or version.")
	listCmd.MarkFlagRequired("type")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
