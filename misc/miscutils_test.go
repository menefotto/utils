package misc

import (
	"testing"

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
