package main

import "fmt"

// declaring the poiner function and we will pass the argument in this function
func pointerFuntion(a *int) {
	// dereferencing
	*a = 578
}
func main() {
	var p = 100
	fmt.Println("Before the function call value of the (P)=", p)

	// taking a pointer variable and assigning the address of p to it
	var pa *int = &p
	// Calling the function by passing pointer to function
	pointerFuntion(pa)

	// After calling the function
	fmt.Println("after calling the function (P) =", p)
}
