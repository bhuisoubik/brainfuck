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
