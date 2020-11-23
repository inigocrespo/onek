package main

import (
	"bytes"
	"fmt"
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

var fd = (int)(os.Stdin.Fd())

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

type normal struct {
	editor *editor
}

type command struct {
	editor  *editor
	command string
	cx, cy  int
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

	var b []byte = make([]byte, 1)
	for {
		os.Stdin.Read(b)
		editor.mode.input(b[0])
	}
}

func (m normal) enter() {
	m.editor.status = "Normal mode"
	m.editor.refresh()
}

func (m command) enter() {
	m.cx = m.editor.cx
	m.cy = m.editor.cy
	m.editor.status = ":"
	m.editor.cx = 1
	m.editor.cy = m.editor.rows
	m.editor.refresh()
}

func (m insert) enter() {
	m.editor.status = "Insert mode"
	m.editor.refresh()
}

func (m normal) leave() {
}

func (m command) leave() {
	m.editor.cx = m.cx
	m.editor.cy = m.cy
}

func (m insert) leave() {
}

func (m normal) input(b byte) {
	switch b {
	case colons:
		m.editor.mode.leave()
		m.editor.mode = m.editor.command
		m.editor.mode.enter()
		break
	case i:
		m.editor.mode.leave()
		m.editor.mode = m.editor.insert
		m.editor.mode.enter()
		break
	case h:
		if m.editor.cx > 0 {
			m.editor.cx = m.editor.cx - 1
			m.editor.refresh()
		}
		break
	case j:
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
		m.editor.mode.leave()
		m.editor.mode = m.editor.normal
		m.editor.mode.enter()
		break
	case enter:
		m.apply()
		break
	default:
		buff := bytes.NewBufferString(m.editor.status)
		buff.WriteByte(b)
		m.editor.status = buff.String()
		m.editor.cx = m.editor.cx + 1
		m.editor.refresh()
	}
}

func (m command) apply() {
	switch m.editor.status {
	case ":":
		m.editor.mode.leave()
		m.editor.mode = m.editor.normal
		m.editor.mode.enter()
		break
	case ":q":
		os.Exit(0)
		break
	default:
	}
}

func (m insert) input(b byte) {
	switch b {
	case esc:
		m.editor.mode.leave()
		m.editor.mode = m.editor.normal
		m.editor.mode.enter()
		break
	case colons:
		return
	default:
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
