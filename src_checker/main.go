package main

import (
	"fmt"
	"os"
	ps "push_swap/pkg"
)

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		return
	}

	stack, err := ps.ConvertArgs(args[0])
	if err != nil {
		fmt.Println("Error")
	} else {
		ps.CheckerProcess(stack)
	}
}
