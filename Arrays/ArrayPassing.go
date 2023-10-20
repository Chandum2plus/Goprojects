package main

import (
	"fmt"
)

// How to pass an Array as an argument in the function
func passArray(a [7]int, size int) int {
	var sum, div int
	for i := 0; i < size; i++ {
		sum += a[i]
	}
	div = sum / size
	fmt.Println("value =", div)
	fmt.Println("size =", size)
	return div
}
func main() {
	var arr = [...]int{49, 60, 40, 34, 56, 78, 90}
	Res := passArray(arr, 1)
	fmt.Print("Result =", Res)
	fmt.Println("\nArray =", arr)
}
