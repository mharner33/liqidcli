package liqtopo

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/mharner33/liqidcli/lapi/liqutil"
	"github.com/tidwall/gjson"
)

const v25num = 9
const separator = "---------------------------------------------------"

func SetupClient(path string) {

}

//Get the current fabric ID
func GetFabID(fid string) string {
	body := liqutil.CallAPI(fid)
	fabID := gjson.Get(string(body), "response.data.0")

	return fabID.String()
}

//Get the counts of each device type
func GetDevicCnt(dPath string) {
	resp := liqutil.CallAPI(dPath)
	results := gjson.GetMany(string(resp), "response.data.0.comp_cnt",
		"response.data.0.targ_cnt",
		"response.data.0.link_cnt",
		"response.data.0.gpu_cnt",
		"response.data.0.plx_cnt",
		"response.data.0.fpga_cnt")

	fmt.Println("Current Device Counts")
	fmt.Println(separator)

	fmt.Printf("Servers: %-10v\n", results[0].Int())
	fmt.Printf("SSDs: %-10v\n", results[1].Int())
	fmt.Printf("NICs: %-10v\n", results[2].Int())
	fmt.Printf("GPUs: %-10v\n", results[3].Int())
	fmt.Printf("PLX: %-10v\n", results[4].Int())
	fmt.Printf("FPGA: %-10v\n", results[5].Int())
}

func GetGroupResource(grpPath string) {
	var gRes liqutil.GroupResource
	body := liqutil.CallAPI(grpPath)

	err := json.Unmarshal(body, &gRes)
	if err != nil {
		fmt.Println("Unable to unmarshal group resource data.")
		os.Exit(1)
	}
	fmt.Printf("\tCPU Count: %-10v\n", gRes.Response.Data[0].CPUCount)
	fmt.Printf("\tGPU Count: %-10v\n", gRes.Response.Data[0].GpuCount)
	fmt.Printf("\tSSD Count: %-10v\n", gRes.Response.Data[0].StorageDriveCount)

}

//Get the current groups and ID's
func GetGroup(grpPath, fid string) {
	//Create query string
	var grp liqutil.GroupList
	qstring := "?parameters=grp_id%3D" + fid
	//fmt.Println(grpPath + qstring)
	body := liqutil.CallAPI(grpPath + qstring)

	err := json.Unmarshal(body, &grp)
	if err != nil {
		fmt.Println("Error unmarshalling group data.")
		os.Exit(1)
	}

	for _, g := range grp.Response.Data {
		fmt.Printf("Group ID: %-10v", g.GrpID)
		fmt.Printf("Group Name: %v\n", g.GroupName)
		GetGroupResource(grpPath + "/details/" + strconv.Itoa(g.GrpID))
	}

}

//List the current machines and their ID's  Will probably need to return a map?
func GetMachines(resp []byte) {

}
func LiqidVersion(path string) {
	//fmt.Printf("In LiqidVersion - %s", path)
	body := liqutil.CallAPI(path)
	var r25 liqutil.Version25
	fieldCnt := gjson.Get(string(body), "response.data.#").Int() //Counts the number of items in the returned array - 9 means 2.5
	if int(fieldCnt) == v25num {                                 //this means it is a 2.5 data map
		err := json.Unmarshal(body, &r25)
		if err != nil {
			fmt.Println("Error unmarshalling 2.5 data!")
		}
		fmt.Println("Liqid Version Information")
		fmt.Println(separator)
		for i := 0; i < int(fieldCnt); i++ {
			fmt.Printf("Component: %-10v  Version: %v\n", r25.Response.Data[i].Component, r25.Response.Data[i].Version)

		}

	} else { //Else we are at a 2.4 code release
		fmt.Println("This is pre-2.5 release of code")
	}
}
func ListSwitch(liqidBase, path string) {
	//Call the propper list methods based on entered command line input
	switch path {
	case "all":
		fmt.Println("List all resources - still working on it")
	case "group":
		//Get the fabric ID so we can list the groups
		fabid := GetFabID(liqidBase + "fabric/id")
		GetGroup(liqidBase+path, fabid)
	case "resource":
		fmt.Println("List resources - still working on it")
	case "machine":
		fmt.Println("List machines - still working on it")
	case "version":
		LiqidVersion(liqidBase + path)
	default:
		fmt.Printf("Bad argument supplied to list command: %s\n", path)
	}

}
