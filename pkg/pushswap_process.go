package push_swap

import (
	"errors"
	"strconv"
	"strings"
)

// Convert a string to valid array of integers
func ConvertArgs(arg string) ([]int, error) {
	var intArr []int
	list := strings.Split(arg, " ")
	for _, nb := range list {
		tmp, err := strconv.Atoi(nb)
		if err != nil {
			return nil, errors.New("error not a number")
		}
		for i := range intArr {
			if tmp == intArr[i] {
				return nil, errors.New("error duplicate")
			}
		}
		intArr = append(intArr, tmp)
	}
	return intArr, nil
}

// Check if array is sorted
func IsSorted(stack []int) bool {
	for i := 0; i < len(stack)-1; i++ {
		if stack[i] > stack[i+1] {
			return false
		}
	}
	return true
}

// Search smallest integer in stack
func FindSmallest(stack []int) int {
	smallest := stack[0]
	for i := range stack {
		if smallest > stack[i] {
			smallest = stack[i]
		}
	}
	return smallest
}

// Give index of N in a stack pass through parameters
func GetIndex(stack []int, n int) int {
	var index int
	for i := range stack {
		if stack[i] == n {
			index = i
		}
	}
	return index
}

// Sort with Push swap actions to stack of length 3
func ShortSort(stack []int) []int {
	smallest := FindSmallest(stack)
	index := GetIndex(stack, smallest)
	switch {
	case index == 0 && stack[1] > stack[2]:
		Sa(stack)
		Ra(stack)

	case index == 1 && stack[0] > stack[2]:
		Ra(stack)

	case index == 1 && stack[0] < stack[2]:
		Sa(stack)

	case index == 2 && stack[0] > stack[1]:
		Sa(stack)
		Rra(stack)

	case index == 2 && stack[0] < stack[1]:
		Rra(stack)

	}
	return stack
}

// Push Swap function who sorts the stack according to the given actions
func Process(a []int) {
	var b []int

	if IsSorted(a) {
		return
	}
	// fmt.Printf("### START a :  %d  | b : %d\n", a, b)

	smallest := FindSmallest(a)

	if a[0] > smallest {
		Sa(a)
	}

find_smallest:
	if len(a) == 3 {
		a = ShortSort(a)
		for len(b) != 0 {
			a, b = Pa(a, b)
		}
		// fmt.Printf("### END a :  %d  | b : %d\n", a, b)
	} else {
		smallest = FindSmallest(a)
		if a[0] == smallest {
			a, b = Pb(a, b)
		} else {
			for a[0] != smallest {
				index := GetIndex(a, smallest)
				if len(a)/2 > index {
					Ra(a)
				} else {
					Rra(a)
				}
			}
		}
		goto find_smallest
	}
}
