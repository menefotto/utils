package futils

import (
	"fmt"
	"os"
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

func TestWriteDirAndFile(t *testing.T) {
	var dmode os.FileMode = 0777
	err := WriteFileOrDir("testdir.txt", []byte(""), os.ModeDir|dmode)
	if err != nil {
		t.Fatal(err)
	}

	var fmode os.FileMode = 0666
	err = WriteFileOrDir("testfile.txt", []byte("ciao"), os.ModeDir|fmode)
	if err != nil {
		t.Fatal(err)
	}

	os.Remove("testdir.txt")
}

func TestCopyFile(t *testing.T) {
	s := "testfile.txt"
	d := "copytest.txt"
	err := CopyFile(s, d)
	if err != nil {
		t.Fatal(err)
	}
	os.Remove(s)
}

func TestMoveFile(t *testing.T) {
	s := "copytest.txt"
	d := "movedtest.txt"
	err := MoveFile(s, d)
	if err != nil {
		t.Fatal(err)
	}
	os.Remove(d)
}
