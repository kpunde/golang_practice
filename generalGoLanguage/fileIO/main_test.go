package fileIO

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

type _format struct {
	Id       string `json:"id"`
	Language string `json:"language"`
	Edition  string `json:"edition"`
	Author   string `json:"author"`
}

func TestTestReadJson(t *testing.T) {
	var formatStruct []_format

	f, err := ioutil.ReadFile("./jsonLog.json")
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(f, &formatStruct)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(formatStruct)
}
