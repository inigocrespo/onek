package main

import (
	"fmt"
)

type normal struct {
	editor *editor
}

func (m normal) enter() {
	m.editor.status = "Normal mode"
	m.editor.draw()
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
			m.editor.draw()
		}

		break
	case j:
		if m.editor.cy < len(m.editor.lines)-1 {
			m.editor.cy = m.editor.cy + 1

			if m.editor.rowoffset < m.editor.rows-2 {
				m.editor.rowoffset = m.editor.rowoffset + 1
			}

			if m.editor.cx > len(m.editor.lines[m.editor.cy])-1 {
				m.editor.cx = len(m.editor.lines[m.editor.cy]) - 1
			}
			m.editor.draw()
		}

		break
	case k:
		if m.editor.cy > 0 {
			m.editor.cy = m.editor.cy - 1

			if m.editor.rowoffset > 0 {
				m.editor.rowoffset = m.editor.rowoffset - 1
			}

			if m.editor.cx > len(m.editor.lines[m.editor.cy])-1 {
				m.editor.cx = len(m.editor.lines[m.editor.cy]) - 1
			}
			m.editor.draw()
		}
		break
	case l:
		if m.editor.cx < len(m.editor.lines[m.editor.cy])-1 {
			m.editor.cx = m.editor.cx + 1
			m.editor.draw()
		}

		break
	default:
		fmt.Println("I got the byte", b, "("+string(b)+")")
	}
}
