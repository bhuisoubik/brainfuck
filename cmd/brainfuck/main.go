/*
MIT License

Copyright (c) 2021 Soubik Bhui (@soubikbhuiwk007)

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

*/

package main

import (
	"fmt"
	"os"
	"io/ioutil"
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
		case "-v",  "version":
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
