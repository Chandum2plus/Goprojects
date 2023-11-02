package main

import "fmt"

// call by value
func callByValue(a, b int) int {
	var c int
	c = a
	a = b
	b = c
	return c

}
func main() {
	var x int = 12
	var y int = 21
	fmt.Printf("x = %d\ny= %d", x, y)
	callByValue(x, y)
	fmt.Printf("\nx= %d, y= %d", x, y)
}
