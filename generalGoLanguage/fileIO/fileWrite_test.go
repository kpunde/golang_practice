package fileIO

import "testing"

func TestWriteFile(t *testing.T) {
	if err := writeFile("smallApache.log"); err != nil {
		t.Errorf("Error found in test: %v", err)
	} else {
		t.Logf("Test executed successfully")
	}
}
