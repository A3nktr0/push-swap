package main

import (
	"fmt"
	"os"
	ps "push_swap/pkg"
)

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		// fmt.Println("Invalid arguments number")
		return
	}

	stack, err := ps.ConvertArgs(args[0])
	if err != nil {
		fmt.Println("Error")
	} else {
		ps.Process(stack)
	}
}
