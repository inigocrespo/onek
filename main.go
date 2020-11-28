package main

import (
	"fmt"
	"os"
)

func main() {

	var file string
	if len(os.Args) == 2 {
		file = os.Args[1]
	} else if len(os.Args) > 2 {
		fmt.Println("Usage: ", os.Args[0], "[filename]")
		return
	}

	editor, err := newEditor(file)
	if err != nil {
		panic(err)
	}

	if err := editor.makeRaw(); err != nil {
		panic(err)
	}

	editor.start()
}
