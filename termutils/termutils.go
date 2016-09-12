package termutils

import (
	"os"
	"syscall"
	"unsafe"
)

func GetTermDimensions() (int, int) {
	fd := os.Stdout.Fd()

	var sz winsize
	_, _, _ = syscall.Syscall(syscall.SYS_IOCTL,
		fd, uintptr(syscall.TIOCGWINSZ), uintptr(unsafe.Pointer(&sz)))
	return int(sz.cols), int(sz.rows)
}

type winsize struct {
	rows    uint16
	cols    uint16
	xpixels uint16
	ypixels uint16
}
