package main

import (
	"fmt"
)

// Passing the simple array into the function
func printArray(array []int) {
	var sum int
	for _, value := range array {
		fmt.Print(" ", value)
		sum += value
	}
	fmt.Println("\nsum of the Array =", sum)
}

// main function
func main() {
	var ar = []int{10, 200, 30, 40, 50}
	printArray(ar)
}
