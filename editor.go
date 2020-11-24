package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/ssh/terminal"
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

type mode interface {
	enter()
	leave()
	input(byte)
}

func newEditor() (*editor, error) {
	width, height, err := terminal.GetSize((int)(os.Stdin.Fd()))
	if err != nil {
		return nil, err
	}

	editor := &editor{}
	normal := normal{editor: editor}
	command := command{editor: editor}
	insert := insert{editor: editor}

	editor.normal = normal
	editor.command = command
	editor.insert = insert
	editor.mode = editor.normal
	editor.cols = width
	editor.rows = height
	editor.cx = 0
	editor.cy = 0

	return editor, nil
}

type editor struct {
	state                   *terminal.State
	normal, command, insert mode
	mode                    mode
	rows, cols              int
	cx, cy                  int
	status                  string
}

func (e *editor) makeRaw() error {
	state, err := terminal.MakeRaw((int)(os.Stdin.Fd()))
	if err != nil {
		return err
	}

	e.state = state
	return nil
}

func (e *editor) start() {
	e.mode.enter()

	var b []byte = make([]byte, 1)
	for {
		os.Stdin.Read(b)
		e.mode.input(b[0])
	}
}

func (e *editor) refresh() {
	os.Stdin.Write([]byte("\x1b[2J"))
	for i := 0; i < e.rows; i++ {
		fmt.Print("~\r\n")
	}
	fmt.Print(e.status)
	fmt.Print("\x1b[", e.cy+1, ";", e.cx+1, "H") // print cursor
}

func (e *editor) exit() {
	terminal.Restore((int)(os.Stdin.Fd()), e.state)
	os.Exit(0)
}
