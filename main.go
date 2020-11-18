package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

var fd = (int)(os.Stdin.Fd())

const (
	esc    = 27
	colons = 58
	q      = 113
	h      = 104
	j      = 106
	k      = 107
	l      = 108
)

var clearScreen = []byte("\x1b[2J")

type editor struct {
	normal     mode
	command    mode
	insert     mode
	mode       mode
	rows, cols int
	cx, cy     int
}

type mode interface {
	enter()
	input(byte)
}

type normal struct {
	editor *editor
}

type command struct {
	editor *editor
}

type insert struct {
	editor *editor
}

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

	editor.mode.enter()
	editor.refresh()

	var b []byte = make([]byte, 1)
	for {
		os.Stdin.Read(b)
		editor.mode.input(b[0])
	}
}

func (m normal) enter() {
}

func (m command) enter() {
}

func (m insert) enter() {
}

func (m normal) input(b byte) {
	switch b {
	case q:
		os.Exit(0)
		break
	case colons:
		m.editor.mode = m.editor.command
		m.editor.mode.enter()
		break
	case h:
		if m.editor.cx > 0 {
			m.editor.cx = m.editor.cx - 1
			m.editor.refresh()
		}
		break
	case j:
		fmt.Println(m.editor.rows)
		if m.editor.cy < m.editor.rows-1 {
			m.editor.cy = m.editor.cy + 1
			m.editor.refresh()
		}
		break
	case k:
		if m.editor.cy > 0 {
			m.editor.cy = m.editor.cy - 1
			m.editor.refresh()
		}
		break
	case l:
		if m.editor.cx < m.editor.cols-1 {
			m.editor.cx = m.editor.cx + 1
			m.editor.refresh()
		}
		break
	default:
		fmt.Println("I got the byte", b, "("+string(b)+")")
	}
}

func (m command) input(b byte) {
	switch b {
	case esc:
		m.editor.mode = m.editor.normal
		m.editor.mode.enter()
		break
	case q:
		os.Exit(0)
		break
	default:
		fmt.Println("Normal mode")
	}
}

func (m insert) input(b byte) {
	switch b {
	case esc:
		m.editor.mode = m.editor.normal
		m.editor.mode.enter()
		break
	case colons:
		return
	default:
		fmt.Println("Isert mode")
	}
}

func (e editor) refresh() {
	os.Stdin.Write(clearScreen)
	fmt.Print("\x1b[", e.cy+1, ";", e.cx+1, "H")
}
