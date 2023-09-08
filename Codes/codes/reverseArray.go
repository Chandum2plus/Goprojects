package main

import "fmt"

// How to Reverse the Array
func reverse() {
	arr := [5]int{10, 20, 30, 40, 50}
	fmt.Print("Original - ")
	for i := 0; i < len(arr); i++ {
		fmt.Print(" ", arr[i])
	}
	fmt.Print("\nReverse -")
	for i := 5 - 1; i >= 0; i-- {
		fmt.Print(" ", arr[i])
	}
}
func main() {
	reverse()
}
