package liqutil

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func CallAPI(urlPath string) []byte {
	//Create the http client
	client := &http.Client{}

	req, _ := http.NewRequest("GET", urlPath, nil)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error getting URL: %s", urlPath)
		os.Exit(1)
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body.")
		os.Exit(2)
	}
	return bodyBytes
}
