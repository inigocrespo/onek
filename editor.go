package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

const (
	enter  = 13
	esc    = 27
	colons = 58
	q      = 113
	h      = 104
	i      = 105
	j      = 106
	k      = 107
	l      = 108
)

type mode interface {
	enter()
	leave()
	input(byte)
}

type editor struct {
	filename                string
	lines                   [][]byte
	state                   *terminal.State
	normal, command, insert mode
	mode                    mode
	rows, cols              int
	rowoffset, coloffset    int
	cx, cy                  int
	status                  string
}

func newEditor(filename string) (*editor, error) {
	width, height, err := terminal.GetSize((int)(os.Stdin.Fd()))
	if err != nil {
		return nil, err
	}

	lines, err := readLines(filename)
	if err != nil {
		return nil, err
	}

	editor := &editor{}
	normal := normal{editor: editor}
	command := command{editor: editor}
	insert := insert{editor: editor}

	editor.filename = filename
	editor.lines = lines
	editor.normal = normal
	editor.command = command
	editor.insert = insert
	editor.mode = editor.normal
	editor.cols = width
	editor.rows = height
	editor.rowoffset = 0
	editor.coloffset = 0
	editor.cx = 0
	editor.cy = 0

	return editor, nil
}

func readLines(filename string) ([][]byte, error) {

	var lines [][]byte = make([][]byte, 0)
	if len(filename) > 0 {
		file, err := os.Open(filename)
		if err != nil {
			return nil, err
		}

		reader := bufio.NewReader(file)
		isNewLine := true
		for {
			line, isPrefix, err := reader.ReadLine()
			if err == io.EOF {
				break
			}

			if err != nil {
				return nil, err
			}

			if isNewLine {
				newLine := make([]byte, len(line), len(line))
				copy(newLine, line)
				lines = append(lines, newLine)
			} else {
				newLine := lines[len(lines)-1]
				newLine = append(newLine, line...)
			}

			isNewLine = !isPrefix
		}
	}

	return lines, nil
}

func (e *editor) makeRaw() error {
	state, err := terminal.MakeRaw((int)(os.Stdin.Fd()))
	if err != nil {
		return err
	}

	e.state = state
	return nil
}

func (e *editor) start() {
	e.mode.enter()

	var b []byte = make([]byte, 1)
	for {
		os.Stdin.Read(b)
		e.mode.input(b[0])
	}
}

func (e *editor) draw() {
	os.Stdin.Write([]byte("\x1b[?25l"))
	os.Stdin.Write([]byte("\x1b[2J"))
	os.Stdin.Write([]byte("\x1b[H"))

	for i := 0; i < e.rows-1; i++ {
		if i < len(e.lines) {
			if i != 0 {
				fmt.Print("\r\n")
			}
			b := e.lines[i]
			fmt.Print(string(b))
		} else {
			fmt.Print("\r\n")
		}
	}

	fmt.Print("\r\n")
	fmt.Print(e.status)
	fmt.Print("\x1b[", e.cy+1, ";", e.cx+1, "H")
	os.Stdin.Write([]byte("\x1b[?25h"))
}

func (e *editor) stop() {
	terminal.Restore((int)(os.Stdin.Fd()), e.state)
	os.Exit(0)
}
