package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

const q = 113

var clearScreen = []byte("\x1b[2J")

type editor struct {
	rows int
	cols int
}

var e = editor{}

func main() {
	state, err := terminal.MakeRaw((int)(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	defer terminal.Restore((int)(os.Stdin.Fd()), state)

	width, height, err := terminal.GetSize((int)(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}

	e.cols = width
	e.rows = height

	write(clearScreen)

	fmt.Println(e.cols, " ", e.rows)

	var b []byte = make([]byte, 1)
	for {
		os.Stdin.Read(b)
		switch b[0] {
		case q:
			return
		default:

			fmt.Println("I got the byte", b, "("+string(b)+")")
		}
	}
}

func write(b []byte) {
	os.Stdin.Write(b)
}
