package misc

import (
	"fmt"
	"os"
	"time"

	"github.com/sonic/lib/utils/terminal"
)

func ProgressMsgBuild(msg string) string {
	w, _ := terminal.GetDimensions()

	var newmsg []byte

	if len(msg) > (w - 19) {
		newmsg = []byte(msg)[:w-19]
	}
	if len(msg) < (w - 19) {
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

func GetDate() string {

	year, month, day := time.Now().Date()
	hour, min, sec := time.Now().Clock()

	zone, _ := time.Now().Zone()
	return fmt.Sprintf("%d %s %d, %d:%d:%d %s",
		day, month.String(), year, hour, min, sec, zone)

}
