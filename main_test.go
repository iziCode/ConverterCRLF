package main

import (
	"testing"
)

func TestGetAllFilesFromPath(t *testing.T) {

	countFiles := len(GetAllFilesFromPath("test"))

	if countFiles != 4 {
		t.Errorf("TestForGetAllFilesFromPath failed, "+
			"count files doesn't match. Expected 4, received %d",
			countFiles)
	}
}
