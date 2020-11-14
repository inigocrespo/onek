package main

import (
	"fmt"
	"os"
)

const q = 113

func main() {
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
