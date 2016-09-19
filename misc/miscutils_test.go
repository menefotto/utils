package misc

import (
	"strings"
	"testing"
	"time"

	"github.com/sonic/lib/utils/terminal"
)

func TestMsgBuildLong(t *testing.T) {
	msg := "fal;hkfl;ashdfl;hasdl;fhl;asdhfl;hsadl;kfhlksadhlkfsdahl;fkla;sflk;sdhl;khjfdafkljadfjadfadljflaksjfldksa"
	cutmsg := ProgressMsgBuild(msg)
	w, _ := terminal.GetDimensions()
	if len(cutmsg) > w {
		t.Error("msg string has not been cut to fit terminal")
	}

	ProgressMsgBuild("ciao")
	ProgressPrinter(cutmsg, 1, 3)
	ProgressPrinter("done", 10000, 100)
}

func TestIsRootUser(t *testing.T) {
	yes := IsRootUser()
	if yes {
		t.Error("should not be root user")
	}
}

func TestGetDate(t *testing.T) {
	date := GetDate()
	datesplit := strings.Split(date, " ")

	_, month, _ := time.Now().Date()

	if datesplit[1] != month.String() {
		t.Fatal("Month should not be different")
	}
}
