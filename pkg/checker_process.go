package push_swap

import (
	"bufio"
	"fmt"
	"os"
)

func Operations(action string, stackA, stackB []int) ([]int, []int) {
	switch action {
	case "pa":
		if len(stackB) > 0 {
			stackA, stackB = Pa(stackA, stackB)
		} else {
			fmt.Println("Invalid Command")
			break
		}
	case "pb":
		if len(stackA) > 0 {
			stackA, stackB = Pb(stackA, stackB)
		} else {
			fmt.Println("Invalid Command")
			break
		}
	case "sa":
		Sa(stackA)
	case "sb":
		Sb(stackB)
	case "ss":
		Ss(stackA, stackB)
	case "ra":
		Ra(stackA)
	case "rb":
		Rb(stackB)
	case "rr":
		Rr(stackA, stackB)
	case "rra":
		Rra(stackA)
	case "rrb":
		Rrb(stackB)
	case "rrr":
		Rrr(stackA, stackB)
	default:
		fmt.Println("Invalid Command")
		return nil, nil
	}
	return stackA, stackB
}

func CheckerProcess(a []int) {
	var b []int

	fmt.Printf("### START a: %d | b: %d\n", a, b)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if len(scanner.Text()) != 0 {
			a, b = Operations(scanner.Text(), a, b)
		}
	}
	fmt.Printf("### END a: %d | b: %d\n", a, b)
	if IsSorted(a) && len(b) == 0 {
		fmt.Println("RESULT : OK")
	} else {
		fmt.Println("RESULT : KO")
	}
}
