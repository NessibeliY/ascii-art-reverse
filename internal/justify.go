package internal

import (
	"errors"
	"strings"
	"syscall"
	"unsafe"
)

func TrimCwd(cwd string) string {
	lastIndex := strings.LastIndex(cwd, "ascii-art-reverse")

	return cwd[:lastIndex+len("ascii-art-reverse")]
}

type winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

func getWidth() (uint, error) {
	ws := &winsize{}

	retCode, _, errno := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)))

	if int(retCode) == -1 {
		return 0, errors.New(string(rune(errno)))
	}

	return uint(ws.Col), nil
}

func width(input string) int {
	width := 0

	bufW := 0

	for index := 0; index < len(input); index++ {
		if rune(input[index]) == '\n' {
			if bufW > width {
				width = bufW
			}

			bufW = 0

		}

		bufW++
	}

	return width
}
