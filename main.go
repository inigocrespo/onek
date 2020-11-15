package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

var fd = (int)(os.Stdin.Fd())

const (
	q = 113
	h = 104
	j = 106
	k = 107
	l = 108
)

var clearScreen = []byte("\x1b[2J")

type editor struct {
	rows, cols int
	cx, cy     int
}

var e = editor{}

func main() {
	state, err := terminal.MakeRaw(fd)
	if err != nil {
		panic(err)
	}
	defer terminal.Restore(fd, state)

	width, height, err := terminal.GetSize(fd)
	if err != nil {
		panic(err)
	}

	e.cols = width
	e.rows = height
	e.cx = 0
	e.cy = 0

	write(clearScreen)
	refresh()

	fmt.Println(e.cols, " ", e.rows)

	var b []byte = make([]byte, 1)
	for {
		os.Stdin.Read(b)
		switch b[0] {
		case q:
			return
		case h:
			if e.cx > 0 {
				e.cx = e.cx - 1
				refresh()
			}
			break
		case j:
			if e.cy < e.rows-1 {
				e.cy = e.cy + 1
				refresh()
			}
			break
		case k:
			if e.cy > 0 {
				e.cy = e.cy - 1
				refresh()
			}
			break
		case l:
			if e.cx < e.cols-1 {
				e.cx = e.cx + 1
				refresh()
			}
			break
		default:
			fmt.Println("I got the byte", b, "("+string(b)+")")
		}
	}
}

func refresh() {
	write(clearScreen)
	fmt.Print("\x1b[", e.cy+1, ";", e.cx+1, "H")
}

func write(b []byte) {
	os.Stdin.Write(b)
}
