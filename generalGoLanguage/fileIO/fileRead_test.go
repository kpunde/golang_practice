package fileIO

import "testing"

func TestReadFullFileInMem(t *testing.T) {
	err := readFullFileInMem("./apache.log")
	if err != nil {
		t.Errorf("Error found in test: %v", err)
	} else {
		t.Logf("Test executed successfully")
	}
}

func TestReadFileInBuffer(t *testing.T) {
	if err := readFileInBuffer("./apache.log"); err != nil {
		t.Errorf("Error found in test: %v", err)
	} else {
		t.Logf("Test executed successfully")
	}
}

func TestReadFileByLine(t *testing.T) {
	if err := readFileByLine("./test"); err != nil {
		t.Errorf("Error found in test: %v", err)
	} else {
		t.Logf("Test executed successfully")
	}
}
