package fileIO

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func readFullFileInMem(fileName string) error {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	return err
}

func readFileInBuffer(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer func() {
		if err = file.Close(); err != nil {
			panic(err)
		}
	}()

	reader := bufio.NewReader(file)
	b := make([]byte, 5)

	for {
		data, err := reader.Read(b)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		fmt.Print(string(b[0:data]))
	}
}

func readFileByLine(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}

	reader := bufio.NewScanner(file)
	for reader.Scan() {
		fmt.Println(string(reader.Bytes()))
	}
	return nil
}
