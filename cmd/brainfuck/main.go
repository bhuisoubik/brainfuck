package main

import (
	"fmt"
	"io/ioutil"
	"os"
	
		"github.com/soubikbhuiwk007/brainfuck/bf"

)

var version = "1.0.2"

func error() {
	fmt.Println("brainfuck: fatal error: file not found\ntry running 'brainfuck -h'")
	os.Exit(1)
}

func help() {
	fmt.Println(`
Brainfuck

Usage:
	brainfuck <command> || <file>

Command:
	-v, version	:	Print Version
	-m <file>	:	Memory view
	-h, help	:	Help (this)`)
}

func readfile(path string) string {
	data, err := ioutil.ReadFile(path)

	if err != nil {
		error()
	}

	return string(data)
}

func main() {
	args := os.Args

	if len(args) < 2 {
		error()
	} else {
		switch args[1] {
		case "-v", "version":
			fmt.Printf("brainfuck version: %v\n", version)
		case "-h", "help":
			help()
		case "-m":
			if len(args) < 3 {
				error()
			} else {
				bf.MemoryView(readfile(args[2]))
			}
		default:
			bf.Parse(readfile(args[1]))
		}
	}
}
