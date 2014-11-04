package main

import (
	"fmt"
	"io/ioutil"
	"bytes"
	"os"
)

func emitGenerator(line []byte) {
	fmt.Print("fmt.Print(\"")
	for _, b := range line {
		switch b {
		case '\n':
			fmt.Print("\\n")
		case '\t':
			fmt.Print("\\t")
		case '"':
			fmt.Print("\\\"")
		case '\\':
			fmt.Print("\\\\")
		default:
			c := 0
			for c != 1 {
				c, _ = os.Stdout.Write([]byte{b})
			}
		}
	}
	fmt.Print("\")\n")
}

func emitCode(line []byte) {
	if len(line) > 0 && line[0] == ' ' {
		line = line[1:]
	}
	fmt.Println(string(line))
}

func emitExpression(expr []byte) {
	fmt.Print("fmt.Printf(\"%v\", ")
	fmt.Print(string(expr))
	fmt.Print(")\n")
}

func whitespaceOnly(line []byte) bool {
	for _, b := range line {
		if b != ' ' && b != '\t' {
			return false
		}
	}
	
	return true
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <filename>\n", os.Args[0])
		os.Exit(1)
	}
	
	file, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	for len(file) > 0 {
		idx1 := bytes.Index(file, []byte{'\n'})
		idx2 := bytes.Index(file, []byte{'/', '*'})
		idx3 := bytes.Index(file, []byte{'/', '/'})

		switch {
		case idx1 != -1 && (idx2 == -1 || idx1 < idx2) && (idx3 == -1 || idx1 < idx3):
			if idx1 > 0 {
				emitGenerator(file[:idx1+1])
			}
			file = file[idx1+1:]
			
		case idx2 != -1 && (idx1 == -1 || idx2 < idx1) && (idx3 == -1 || idx2 < idx3):
			emitGenerator(file[:idx2])
			
			file = file[idx2+2:]
			idx4 := bytes.Index(file, []byte{'*', '/'})
			if idx4 == -1 {
				panic("bad input file")
			}
			
			emitExpression(file[:idx4])
			file = file[idx4+2:]
			
		case idx3 != -1:
			if idx1 == -1 {
				panic("bad input file")
			}
			
			if !whitespaceOnly(file[:idx3]) {
				emitGenerator(file[:idx3])
			}
			emitCode(file[idx3+2:idx1])
			file = file[idx1+1:] 

		default:
			emitGenerator(file)
			file = nil
		}
	}
}
