package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

const q = 113

func main() {
	state, err := terminal.MakeRaw((int)(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	defer terminal.Restore((int)(os.Stdin.Fd()), state)

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
