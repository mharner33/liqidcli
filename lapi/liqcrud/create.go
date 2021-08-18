package liqcrud

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mharner33/liqidcli/lapi/liqutil"
)

func CreateGroup(gPath string, gInfo liqutil.CreateGrpStruct) {

	jsonVal, _ := json.Marshal(gInfo)
	_, err := http.Post(gPath, "application/json", bytes.NewBuffer(jsonVal))
	if err != nil {
		fmt.Printf("The HTTP post failed with error %s\n", err)
	}

}
