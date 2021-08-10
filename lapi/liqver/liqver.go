package liqver

import (
	"encoding/json"
	"fmt"

	"github.com/tidwall/gjson"
)

const v24num = 4
const v25num = 9

var ver24fields = []string{"liqcfgd-version", "liqmond-version", "liqmgt-version", "ui-version"}

type Response24 struct {
	Response struct {
		Data []struct {
			LiqcfgdVersion struct {
				Major  int `json:"major"`
				Minor  int `json:"minor"`
				Patch  int `json:"patch"`
				Number int `json:"number"`
			} `json:"liqcfgd-version"`
			LiqmondVersion struct {
				Major  int `json:"major"`
				Minor  int `json:"minor"`
				Patch  int `json:"patch"`
				Number int `json:"number"`
			} `json:"liqmond-version"`
			LiqmgtVersion struct {
				Major  int `json:"major"`
				Minor  int `json:"minor"`
				Patch  int `json:"patch"`
				Number int `json:"number"`
			} `json:"liqmgt-version"`
			UIVersion struct {
				Major  int `json:"major"`
				Minor  int `json:"minor"`
				Patch  int `json:"patch"`
				Number int `json:"number"`
			} `json:"ui-version"`
		} `json:"data"`
		Errors interface{} `json:"errors"`
		Code   int         `json:"code"`
	} `json:"response"`
}

type Response25 struct {
	Response struct {
		Data []struct {
			Component      string `json:"component"`
			Branch         string `json:"branch"`
			Changeset      string `json:"changeset"`
			Date           string `json:"date"` //was time.time
			Version        string `json:"version"`
			ChangesetShort string `json:"changeset_short"`
			DateShort      string `json:"date_short"`
		} `json:"data"`
		Errors interface{} `json:"errors"`
		Code   int         `json:"code"`
	} `json:"response"`
}

func LiqidVersion(body []byte) {
	var r25 Response25
	fieldCnt := gjson.Get(string(body), "response.data.#").Int() //Counts the number of items in the returned array - 9 means 2.5
	if int(fieldCnt) == v25num {                                 //this means it is a 2.5 data map
		err := json.Unmarshal(body, &r25)
		if err != nil {
			fmt.Println("Error unmarshalling 2.5 data!")
		}
		for i := 0; i < int(fieldCnt); i++ {
			fmt.Printf("Component: %-10v  Version: %v\n", r25.Response.Data[i].Component, r25.Response.Data[i].Version)

		}

	} else { //Else we are at a 2.4 code release
		gmaj := gjson.Get(string(body), "response.data.#.liqcfgd-version.major")
		gmin := gjson.Get(string(body), "response.data.#.liqcfgd-version.minor")
		gnum := gjson.Get(string(body), "response.data.#.liqcfgd-version.number")
		for i := 0; i < v24num; i++ {
			fmt.Printf("Component: %-10v  Version: %v.%v.%v\n", ver24fields[i], gmaj.Int(), gmin.String(), gnum.String())
		}
	}
}
