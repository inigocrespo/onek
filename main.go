package main

import (
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

var fd = (int)(os.Stdin.Fd())

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
