package miscutils

import (
	"fmt"
	"os"

	"github.com/sonic/lib/utils/termutils"
)

func ProgressMsgBuild(msg string) string {
	w, _ := termutils.GetDimensions()

	var newmsg []byte

	switch {
	case len(msg) > (w - 19):
		newmsg = []byte(msg)[:w-19]
	case len(msg) < (w - 19):
		newmsg = []byte(msg)
	}

	spacen := w - (len(msg) + 9)
	spaces := []byte(" ")

	for i := 0; i < spacen; i++ {
		spaces = append(spaces, []byte(" ")...)
	}

	return string(newmsg) + string(spaces)
}

func ProgressPrinter(msg string, tot, percent int64) {
	if tot/percent == 100 {
		fmt.Fprintf(os.Stdout, "%s%d%s\n", msg, 100, "%")
	} else {
		fmt.Fprintf(os.Stdout, "%s%d%s\r", msg, tot/percent, "%")
	}
}

func IsRootUser() bool {
	if os.Geteuid() != 0 {
		return false
	}
	return true
}
