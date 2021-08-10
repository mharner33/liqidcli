package liqtopo

import (
	"fmt"

	"github.com/tidwall/gjson"
)

//structure to keep track of the number of the various components in the fabric
type DevCounts struct {
	Servers int64
	Ssd     int64
	Nic     int64
	Gpu     int64
	Plx     int64
	Fpga    int64
}

var DCount = make(map[string]int64)

func GetFabID(resp []byte) string {
	fabID := gjson.Get(string(resp), "response.data.0")

	return fabID.String()
}

func GetDevicCnt(resp []byte) {

	DCount, ok := gjson.Parse(string(resp)).Value().(map[string]int64)
	if !ok {
		fmt.Println("Error - not a map in GetDeviceCnt")
	}
	fmt.Println("In GetDeviceCnt", DCount["targ_cnt"])
	// results := gjson.GetMany(string(resp), "response.data.0.comp_cnt",
	// 	"response.data.0.targ_cnt",
	// 	"response.data.0.link_cnt",
	// 	"response.data.0.gpu_cnt",
	// 	"response.data.0.plx_cnt",
	// 	"response.data.0.fpga_cnt")

	// dcount.Servers = results[0].Int()
	// dcount.Ssd = results[1].Int()
	// dcount.Nic = results[2].Int()
	// dcount.Gpu = results[3].Int()
	// dcount.Plx = results[4].Int()
	// dcount.Fpga = results[5].Int()

}

//Get the current groups and ID's - will probably need to return a map?
func GetGroup(resp []byte) {

}

//List the current machines and their ID's  Will probably need to return a map?
func GetMachines(resp []byte) {

}
