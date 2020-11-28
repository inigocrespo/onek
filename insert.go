package main

type insert struct {
	editor *editor
}

func (m insert) enter() {
	m.editor.status = "Insert mode"
	m.editor.draw()
}

func (m insert) leave() {
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
