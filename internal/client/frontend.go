package client

import (
	"syscall"
	"unsafe"
	//tea "github.com/charmbracelet/bubbletea"
)

type winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

func getDimensions() (uint, uint) {
	ws := &winsize{}
	retCode, _, errno := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)))

	if int(retCode) == -1 {
		panic(errno)
	}
	return uint(ws.Row), uint(ws.Col)
}
