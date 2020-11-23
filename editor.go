package main

import (
	"fmt"
	"os"
)

const (
	enter  = 13
	esc    = 27
	colons = 58
	q      = 113
	h      = 104
	i      = 105
	j      = 106
	k      = 107
	l      = 108
)

type editor struct {
	normal, command, insert mode
	mode                    mode
	rows, cols              int
	cx, cy                  int
	status                  string
}

type mode interface {
	enter()
	leave()
	input(byte)
}

func (e *editor) refresh() {
	os.Stdin.Write([]byte("\x1b[2J"))
	for i := 0; i < e.rows; i++ {
		fmt.Print("~\r\n")
	}
	fmt.Print(e.status)
	fmt.Print("\x1b[", e.cy+1, ";", e.cx+1, "H") // print cursor
}
