package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

var (
	ver = "1.0.0"
	mem = make([]byte, 1)
	current = 0
	fPtr int
)

func main() {
	args := os.Args
	mem[0] = 0
	if len(args) < 2 {
		interactiveConsole()
	} else {
		switch args[1] {
		case "-v", "version":
			fmt.Printf("Version: %v\n", ver)
		case "-h", "help" :
			printHelp()
		case "run" :
			data, _ := ioutil.ReadFile(args[2])
			interpret(data)
		case "-m", "memory":
			data, _ := ioutil.ReadFile(args[2])
			interpret(data)
			fmt.Println("\nMemory: ", mem)
		default:
			fmt.Printf("Command <%v> not found..\n", args[1])
			printHelp()
		}
	}
}

func printHelp() { fmt.Println("Usage : \n\tbrainfuck <command>\nCommands :\n\t run [filename]\t\t:\tTo Interpret a Brainfuck (*.bf) file\n\t -v, version\t\t:\tTo Print the Current Version\n\t -m, memory [filename]\t:\tTo Run & Print the Memory usage of a Brainfuck (*.bf) file\n\t -h, help\t\t:\tFor Help")}

func interactiveConsole() {
	fmt.Printf("Interactive Brainfuck (v%v)\nType exit to Exit", ver)
	src,inp, exit := "","", false
	for !exit {
		fmt.Print("\nbrainfuck >> ")
		fmt.Scanln(&inp)
		if inp == "exit" {
			exit = true
		} else {
			src += inp
			interpret([]byte(src))
		}
	}
}

func interpret(src []byte) {
	for i := 0; i < len(src); i++ {
		if src[i] == '>' {
			if len(mem) - 1 <= current {
				current++
				mem = append(mem, 0)
			} else {
				current++
			}
		} else if src[i] == '<' {
			if current >= 0 {
				current--
			} else {
				fmt.Println("Error: Memory Pointer can't be less than 0")
				break
			}
		} else if src[i] == '+' {
			if mem[current] == 255 {
				mem[current] = 0
			} else {
				mem[current]++
			}
		} else if src[i] == '-' {
			if mem[current] == 0 {
				mem[current] = 255
			} else {
				mem[current]--
			}
		} else if src[i] == '.' {
			fmt.Print(string(mem[current]))
		} else if src[i] == ',' {
			inp, _ := bufio.NewReader(os.Stdin).ReadString('\n')
			mem[current] = []byte(inp)[0]
		} else if src[i] == '[' {
			fPtr = i
		} else if src[i] == ']' {
			if mem[current] != 0 {
				i = fPtr
			}
		}
	}
}