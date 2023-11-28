package push_swap

import "fmt"

// push the top first element of stack b to stack a
func Pa(a, b []int) ([]int, []int) {
	tmp := b[0]
	b = b[1:]
	a = append(a, tmp)
	tmp = a[len(a)-1]
	for i := len(a) - 1; i > 0; i-- {
		a[i] = a[i-1]
	}
	a[0] = tmp
	fmt.Println("pa")
	return a, b
}

// push the top first element of stack a to stack b
func Pb(a, b []int) ([]int, []int) {
	tmp := a[0]
	a = a[1:]
	b = append(b, tmp)
	tmp = b[len(b)-1]
	for i := len(b) - 1; i > 0; i-- {
		b[i] = b[i-1]
	}
	b[0] = tmp
	fmt.Println("pb")
	return a, b
}

// swap first 2 elements of stack a
func Sa(a []int) {
	a[0], a[1] = a[1], a[0]
	fmt.Println("sa")
}

// swap first 2 elements of stack b
func Sb(b []int) {
	b[0], b[1] = b[1], b[0]
	fmt.Println("sb")
}

// execute sa and sb
func Ss(a, b []int) {
	Sa(a)
	Sb(b)
	fmt.Println("ss")
}

// rotate stack a (shift up all elements of stack a by 1, the first element becomes the last one)
func Ra(a []int) {
	tmp := a[0]
	for i := 0; i < len(a)-1; i++ {
		a[i] = a[i+1]
	}
	a[len(a)-1] = tmp
	fmt.Println("ra")
}

// rotate stack b
func Rb(b []int) {
	tmp := b[0]
	for i := 0; i < len(b)-1; i++ {
		b[i] = b[i+1]
	}
	b[len(b)-1] = tmp
	fmt.Println("rb")
}

// execute ra and rb
func Rr(a, b []int) {
	Ra(a)
	Rb(b)
	fmt.Println("rr")
}

// reverse rotate a (shift down all elements of stack a by 1, the last element becomes the first one)
func Rra(a []int) {
	tmp := a[len(a)-1]
	for i := len(a) - 1; i > 0; i-- {
		a[i] = a[i-1]
	}
	a[0] = tmp
	fmt.Println("rra")
}

// reverse rotate b
func Rrb(b []int) {
	tmp := b[len(b)-1]
	for i := len(b) - 1; i > 0; i-- {
		b[i] = b[i-1]
	}
	b[0] = tmp
	fmt.Println("rrb")

}

// execute rra and rrb
func Rrr(a, b []int) {
	Rra(a)
	Rrb(b)
	fmt.Println("rrr")
}
