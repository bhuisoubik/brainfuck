package bf

import (
	"fmt"
)

var mem = make([]byte, 1)

func Parse(src string) {
	current, fPtr := 0, 0
	for i := 0; i < len(src); i++ {
		if src[i] == '>' {
			if len(mem)-1 <= current {
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
			mem[current] = input()
		} else if src[i] == '[' {
			fPtr = i
		} else if src[i] == ']' {
			if mem[current] != 0 {
				i = fPtr
			}
		}
	}
}

func MemoryView(src string) {
	Parse(src)
	fmt.Printf("\nMemory: %v\n", mem)
}
