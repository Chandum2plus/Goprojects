package main

import (
	"fmt"
)

// How to find the even and odd number of a slice and sum of the even and odd slice

func evenOdd() {

	slic := []int{2, 4, 5, 6, 7, 7}
	even := 0
	odd := 0
	evensum := 0
	Oddsum := 0
	fmt.Println("Even Number of the slice")
	for i := 0; i < len(slic); i++ {
		if slic[i]%2 == 0 {
			fmt.Println(slic[i])
			even++
			evensum += slic[i]
		}
	}
	fmt.Println("Total Number of the Even =", even)
	fmt.Println("Sum of the Even slice - ", evensum)
	fmt.Println()
	fmt.Println("Number of the Odd slices")
	for i := 0; i < len(slic); i++ {
		if slic[i]%2 != 0 {
			fmt.Println(slic[i])
			odd++
			Oddsum += slic[i]
		}
	}
	fmt.Println("Total Number of the Odd  = ", odd)
	fmt.Println("Sum of the Odd slice - ", Oddsum)
}
func main() {
	evenOdd()
}
