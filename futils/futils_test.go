package futils

import (
	"fmt"
	"log/syslog"
	"os"
	"testing"
	"time"

	"github.com/sonic/lib/syslogger"
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

func TestFileMover(t *testing.T) {
	s := "/tmp/"
	d := "/var/cache/pacman/pkg/"
	filename := "tar-1.29-1-x86_64.pkg.tar.xz"

	err := CopyFile(d+filename, s+filename)
	if err != nil {
		t.Fatal(err)
	}

	mover := NewFileMover(syslogger.NewLogger("testmover", syslog.LOG_ERR))
	mover.Send(s+filename, d+filename)
	time.Sleep(time.Second * 2)
	mover.Close()
}

func TestFileMoverManyMoreThen32(t *testing.T) {
	s := "/tmp/"
	d := "/var/cache/pacman/pkg/"
	filename := "tar-1.29-1-x86_64.pkg.tar.xz"

	err := CopyFile(d+filename, s+filename)
	if err != nil {
		t.Fatal(err)
	}

	mover := NewFileMover(syslogger.NewLogger("testmover", syslog.LOG_ERR))
	for i := 35; i > 0; i-- {
		mover.Send(s+filename, d+filename)
	}
	time.Sleep(time.Second * 0)
	mover.Close()
}

func TestUniquePath(t *testing.T) {
	slice := []string{"/usr/test/", "/usr/test"}
	res := UniquePaths(slice)
	if len(res) != 1 {
		t.Error("shuold be only one path")
	}

}
