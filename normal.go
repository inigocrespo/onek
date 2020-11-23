package main

import (
	"fmt"
)

type normal struct {
	editor *editor
}

func (m normal) enter() {
	m.editor.status = "Normal mode"
	m.editor.refresh()
}

func (m normal) leave() {
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
