package liqtopo

import (
	"encoding/json"

	"github.com/mharner33/liqidcli/lapi/liqutil"
	"github.com/tidwall/gjson"
)

func SetupClient(path string) {

}

//Get the current fabric ID
func GetFabID(resp []byte) string {
	fabID := gjson.Get(string(resp), "response.data.0")

	return fabID.String()
}

//Get the counts of each device type
func GetDevicCnt(resp []byte, dcnt *liqutil.DevCounts) {
	results := gjson.GetMany(string(resp), "response.data.0.comp_cnt",
		"response.data.0.targ_cnt",
		"response.data.0.link_cnt",
		"response.data.0.gpu_cnt",
		"response.data.0.plx_cnt",
		"response.data.0.fpga_cnt")

	dcnt.Servers = results[0].Int()
	dcnt.Ssd = results[1].Int()
	dcnt.Nic = results[2].Int()
	dcnt.Gpu = results[3].Int()
	dcnt.Plx = results[4].Int()
	dcnt.Fpga = results[5].Int()

}

//Get the current groups and ID's
func GetGroup(resp []byte, garr []liqutil.Groups) {
	json.Unmarshal(resp, &garr)
}

//List the current machines and their ID's  Will probably need to return a map?
func GetMachines(resp []byte) {

}

func GetAll() {
	//Get the fabric ID

}
