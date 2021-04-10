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
	"bufio"
	"fmt"
	"github.com/soubikbhuiwk007/brainfuck/bf"
	"os"
)

var version = "1.0.0"

func ibf() {
	fmt.Printf("Interactive Brainfuck (v%v)\nType exit to Exit", version)
	src, inp, exit := "", "", false
	for !exit {
		fmt.Print("\nibf >> ")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			inp = scanner.Text()
			src += inp
		}
		if inp == "exit" {
			exit = true
		} else {
			src += inp
			bf.Parse(src)
		}
	}
}

func error() {
	fmt.Println("brainfuck: fatal error: file not found\ntry running 'brainfuck -h'")
	os.Exit(1)
}

func main() {
	args := os.Args

	if len(args) >= 2 {
		if args[1] == "-v" || args[1] == "version" {
			fmt.Printf("ibf version: %v\n", version)
		} else {
			error()
		}
	} else {
		ibf()
	}
}
