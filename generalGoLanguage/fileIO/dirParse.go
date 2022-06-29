package fileIO

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func parseDir(dirName string, intend string) {
	files, err := ioutil.ReadDir(dirName)
	if err != nil {
		fmt.Println(err)
	}
	for _, f := range files {
		isFile := "f"

		if f.IsDir() {
			isFile = "d"
		}
		fmt.Printf("%v--%v-- Name: %v \n", intend, isFile, f.Name())
		if f.IsDir() {
			path := filepath.Join(dirName, f.Name())
			//fmt.Println(path)
			intend += "    "
			parseDir(path, intend)
		}
	}
}
