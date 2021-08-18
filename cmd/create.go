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
	"strconv"
	"strings"

	"github.com/mharner33/liqidcli/lapi/liqcrud"
	"github.com/mharner33/liqidcli/lapi/liqtopo"
	"github.com/mharner33/liqidcli/lapi/liqutil"
	"github.com/spf13/cobra"
)

var createType string
var name string
var gid string

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create resources such as groups and machines",
	Long: `This command creates any desired resources such as groups or machines in
	the Liqid system.  
	Valid types:  group <group name>, machine <machine name> <group id>
	
	Some examples include:
	
	liqidcli --ip 10.204.105.38 create --type group --name mygroup
	liqidcli --ip 10.204.106.130 create --type machine mymachine 1`,

	Run: func(cmd *cobra.Command, args []string) {
		createType = strings.ToLower(createType)
		basePath := "http://" + ipAddress + liqutil.ApiPath
		fmt.Println(basePath + createType)
		fabid := liqtopo.GetFabID(basePath + "fabric/id")
		var GrpInfo liqutil.CreateGrpStruct
		switch createType {
		case "group":

			gid := liqutil.GetNext(basePath, "group")

			GrpInfo.FabID, _ = strconv.Atoi(fabid)
			GrpInfo.GrpName = name
			GrpInfo.GrpID = gid
			GrpInfo.PodID = -1
			liqcrud.CreateGroup(basePath+createType, GrpInfo)
		default:

		}

	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().StringVarP(&createType, "type", "t", "", "Display information about the Liqid environment.")
	createCmd.Flags().StringVarP(&name, "name", "n", "Test123", "The name of the group or machine to create.")
	createCmd.Flags().StringVarP(&gid, "gid", "g", "", "Group ID in which to create the machine in.")
	createCmd.MarkFlagRequired("type")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
