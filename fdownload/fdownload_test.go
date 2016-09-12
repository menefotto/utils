package fdownload

import (
	"os"
	"testing"
)

func TestClietInit(t *testing.T) {
	client := clientInit()
	_, err := client.Get("https://www.google.com")
	if err != nil {
		t.Error("Error: ", err)
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

func TestDownloadSequential(t *testing.T) {
	base := "http://archlinux.polymorf.fr/core/os/x86_64/"

	pkgnames := []string{
		"acl-2.2.52-2-x86_64.pkg.tar.xz",
		"automake-1.15-1-any.pkg.tar.xz"}

	for _, pkgname := range pkgnames {
		err := DownloadSingle(base, ".", pkgname)
		if err != nil {
			t.Error(err)
		}
		os.Remove(pkgname)
	}

}
