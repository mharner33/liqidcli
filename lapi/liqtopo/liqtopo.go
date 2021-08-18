package liqtopo

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/mharner33/liqidcli/lapi/liqutil"
)

const separator = "---------------------------------------------------"

//Get the current fabric ID - /liqid/api/v2/fabric/id
func GetFabID(path string) string {
	var fid liqutil.FabricID
	body := liqutil.CallAPI(path)
	err := json.Unmarshal(body, &fid)
	if err != nil {
		fmt.Println("Error parsing fabric id")
		os.Exit(1)
	}

	return strconv.Itoa(fid.Response.Data[0])
}

//Get the counts of each device type
//from /liqid/api/v2/devices/count?fabr_id=<fab_id>
func GetDevicCnt(dPath string) {
	var dCount liqutil.DeviceCount
	body := liqutil.CallAPI(dPath)
	err := json.Unmarshal(body, &dCount)
	if err != nil {
		fmt.Println("Error parsing device count information")
		os.Exit(1)
	}
	fmt.Println("Current Device Counts")
	fmt.Println(separator)

	fmt.Printf("Servers: %-10v\n", dCount.Response.Data[0].CompCnt)
	fmt.Printf("SSDs: %-10v\n", dCount.Response.Data[0].TargCnt)
	fmt.Printf("NICs: %-10v\n", dCount.Response.Data[0].LinkCnt)
	fmt.Printf("GPUs: %-10v\n", dCount.Response.Data[0].GpuCnt)
	fmt.Printf("PLX: %-10v\n", dCount.Response.Data[0].PlxCnt)
	fmt.Printf("FPGA: %-10v\n", dCount.Response.Data[0].FpgaCnt)
}

func GetGroupResource(grpPath string) {
	var gRes liqutil.GroupResource
	fmt.Println(grpPath)
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
func GetGroups(grpPath, qString string) {

	var grp liqutil.GroupList

	//fmt.Println(grpPath + qString)
	body := liqutil.CallAPI(grpPath + qString)

	err := json.Unmarshal(body, &grp)
	if err != nil {
		fmt.Println("Error unmarshalling group data.")
		os.Exit(1)
	}

	for _, g := range grp.Response.Data {
		fmt.Printf("Group ID: %-10v", g.GrpID)
		fmt.Printf("Group Name: %v\n", g.GroupName)
		GetGroupResource(grpPath + "group/details/" + strconv.Itoa(g.GrpID))
	}

}

//List details about a particular machine - used to get machines in a group
func GetMachines() {

}

func GetVersion(path string) {
	//fmt.Printf("In LiqidVersion - %s", path)
	body := liqutil.CallAPI(path)
	var ver liqutil.Version
	err := json.Unmarshal(body, &ver)
	if err != nil {
		fmt.Println("Error parsing version json - this only works with 2.5+ versions.")
		os.Exit(1)
	}
	for _, v := range ver.Response.Data {
		fmt.Printf("Component: %-10v Version: %v\n", v.Component, v.Version)
	}
}
