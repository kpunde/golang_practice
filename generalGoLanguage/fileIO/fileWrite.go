package fileIO

import (
	"bufio"
	"os"
)

func writeFile(fileName string) error {
	inF, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer inF.Close()
	reader := bufio.NewScanner(inF)
	opF, err := os.OpenFile("test", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer opF.Close()
	for reader.Scan() {
		_, err := opF.Write(reader.Bytes())
		_, err = opF.Write([]byte("\n"))
		if err != nil {
			return err
		}
	}

	return nil
}
