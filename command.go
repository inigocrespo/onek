package main

import (
	"bytes"
)

type command struct {
	editor  *editor
	command string
	cx, cy  int
}

func (m command) enter() {
	m.cx = m.editor.cx
	m.cy = m.editor.cy
	m.editor.status = ":"
	m.editor.cx = 1
	m.editor.cy = m.editor.rows
	m.editor.draw()
}

func (m command) leave() {
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
		m.editor.draw()
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
		m.editor.stop()
		break
	default:
	}
}
