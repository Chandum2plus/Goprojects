package main

import "fmt"

// Reverse the Array
func reverse(input []int) []int {
	var output []int
	for i := len(input) - 1; i >= 0; i-- {
		output = append(output, input[i])
	}
	return output

}
func main() {
	var size int
	fmt.Print("Enter the size of the slice ")
	fmt.Scan(&size)

	var sl = make([]int, size)

	fmt.Print("Enter the element - ")
	for i := 0; i < size; i++ {
		fmt.Scan(&sl[i])
	}
	fmt.Println("Reverse = ", reverse(sl))

}
