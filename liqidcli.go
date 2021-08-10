package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	_ "github.com/mharner33/liqidcli/cmd"
	"github.com/mharner33/liqidcli/lapi/liqtopo"
	"github.com/mharner33/liqidcli/lapi/liqutil"
)

var testURL = "http://10.204.105.38:8080/liqid/api/v2/"

//var testPath = "fabric/id"
var testPath = "devices/count"

func main() {
	//Create the http client
	client := &http.Client{}
	//cmd.Execute()

	//Test the fabric ID
	req, _ := http.NewRequest("GET", testURL+testPath, nil)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error getting URL: %s", testURL+testPath)
		os.Exit(1)
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body.")
		os.Exit(2)
	}
	//fid := liqtopo.GetFabID(bodyBytes)

	//fmt.Printf("Fabric ID is: %s\n", fid)
	////////////////////////////////////////////////////
	//Test the device counts
	var dcounts = new(liqutil.DevCounts)

	liqtopo.GetDevicCnt(bodyBytes, dcounts)
	fmt.Printf("Servers: %d\n", dcounts.Servers)
	fmt.Printf("SSD: %d\n", dcounts.Ssd)
	fmt.Printf("GPU: %d\n", dcounts.Gpu)
}
