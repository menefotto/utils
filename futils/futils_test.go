package futils

import (
	"fmt"
	"testing"
)

func TestDirList(t *testing.T) {
	directory := "/home/carlo/Work"
	list, err := DirList(directory)
	if err != nil {
		t.Logf("Error %v\n", err)
	}

	for _, dir := range list {
		fmt.Println("Directory :", dir)
	}
	fmt.Println(list)
}
