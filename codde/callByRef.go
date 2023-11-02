package main

import "fmt"

// Example of the call by  Reference
func Ref(a, b *int) int {
	var c int
	c = *a
	*a = *b
	*b = c
	return c
}
func main() {
	x := 12
	y := 10
	fmt.Printf("x= %d\ny=%d", x, y)
	Ref(&x, &y)
	fmt.Println("\nAFTER")
	fmt.Printf("\nx=%d\ny=%d", x, y)
}
