package fileIO

import (
	"testing"
)

type _format struct {
	Id       string `json:"id"`
	Language string `json:"language"`
	Edition  string `json:"edition"`
	Author   string `json:"author"`
}

func TestTestReadJson(t *testing.T) {
	parseDir("/Users/kartikpunde/go/src/dsa", "")
}
