package fdownload

import (
	"os"
	"testing"

	"github.com/sonic/lib/utils/termutils"
)

func TestClietInit(t *testing.T) {
	client := clientInit()
	_, err := client.Get("https://www.google.com")
	if err != nil {
		t.Error("Error: ", err)
	}

}

func TestClietInitWrong(t *testing.T) {
	client := clientInit()
	_, err := client.Get("ht//www.google.com")
	if err == nil {
		t.Error("Should not be nil")
	}

}
func TestSingleDownload(t *testing.T) {
	base := "http://archlinux.polymorf.fr/core/os/x86_64/"
	pkgname := "bash-4.3.046-1-x86_64.pkg.tar.xz"
	err := DownloadSingle(base, ".", pkgname)
	if err != nil {
		t.Error(err)
	}
	os.Remove(pkgname)
}
func TestSingleDownloadWrong(t *testing.T) {
	base := "http://archlinux.fr/core/os/x86_64/"
	pkgname := "bash-4.3.046-1-x86_64.pkg.tar.xz"
	err := DownloadSingle(base, ".", pkgname)
	if err == nil {
		t.Error("Should not be nil")
	}
}

func TestDownloadMulti(t *testing.T) {
	base := "http://archlinux.polymorf.fr/core/os/x86_64/"

	pkgnames := []string{
		"acl-2.2.52-2-x86_64.pkg.tar.xz",
		"bash-4.3.046-1-x86_64.pkg.tar.xz",
	}

	errchan := DownloadMulti(base, ".", pkgnames)
	for i := 0; i < len(pkgnames); i++ {
		<-errchan
	}

	for _, name := range pkgnames {
		err := os.Remove(name)
		if err != nil {
			//fmt.Println("Err :", err)
		}
	}
}
func TestDownloadMultiWrong(t *testing.T) {
	base := "http://archlinux.polymorf.fr/core/os/x86_64/"

	pkgnames := []string{
		"acls2-2-x86_64.pkg.tar.xz",
		"bass-4.3.046-1-x86_64.pkg.tar.xz",
	}

	errchan := DownloadMulti(base, ".", pkgnames)
	for i := 0; i < len(pkgnames); i++ {
		msg := <-errchan
		if msg == nil {
			t.Error("there should be an error")
		}
	}

}
func TestDownloadSequential(t *testing.T) {
	base := "http://archlinux.polymorf.fr/core/os/x86_64/"

	pkgnames := []string{
		"acl-2.2.52-2-x86_64.pkg.tar.xz",
		"acl-2.2.52-2-x86_64.pkg.tar.xz",
	}

	for _, pkgname := range pkgnames {
		err := DownloadSingle(base, ".", pkgname)
		if err != nil {
			t.Error(err)
		}
		os.Remove(pkgname)
	}

}

func TestMsgBuildLong(t *testing.T) {
	msg := "fal;hkfl;ashdfl;hasdl;fhl;asdhfl;hsadl;kfhlksadhlkfsdahl;fkla;sflk;sdhl;khjfdafkljadfjadfadljflaksjfldksa"
	cutmsg := progressMsgBuild(msg)
	w, _ := termutils.GetDimensions()
	if len(cutmsg) > w {
		t.Error("msg string has not been cut to fit terminal")
	}
}
