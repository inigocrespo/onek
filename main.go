package main

import (
	"os"
)

var fd = (int)(os.Stdin.Fd())

func main() {

	editor, err := newEditor()
	if err != nil {
		panic(err)
	}

	if err := editor.makeRaw(); err != nil {
		panic(err)
	}

	editor.start()
}
