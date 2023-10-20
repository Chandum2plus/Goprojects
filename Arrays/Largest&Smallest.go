package main

import (
	"fmt"
)

// Find the largest and smallest number in Array

func largeSmall() {
	var size, minposition, maxposition, largest, smallest int
	fmt.Print("Enter the size of the Array - ")
	fmt.Scan(&size)
	///	i := 0
	// Declaring the Array -
	var arr = make([]int, size)

	fmt.Print("Enter the Element - ")
	for i := 0; i < size; i++ {
		fmt.Scan(&arr[i])
	}
	largest = arr[0]
	smallest = arr[0]
	for i := 0; i < size; i++ {
		if largest < arr[i] {
			largest = arr[i]
			maxposition = i
		}
		if smallest > arr[i] {
			smallest = arr[i]
			minposition = i
		}
	}
	fmt.Println("largest - ", largest)
	fmt.Println("Max position - ", maxposition)
	fmt.Println("Smallest - ", smallest)
	fmt.Println("Min position - ", minposition)
}

// Print the even number  in index position
func evenIndex() {
	var size int
	fmt.Print("Enter the size of the Array - ")
	fmt.Scan(&size)

	var arr = make([]int, size)

	fmt.Print("Enter the element - ")
	for i := 0; i < size; i++ {

		fmt.Scan(&arr[i])
	}
	i := 0
	fmt.Println("The list of the Array item in even index position -")

	for i = 0; i < size; i++ {
		if arr[i]%2 == 0 {
			fmt.Print(" ", arr[i])
		}
	}

}

// How to print the odd number in index position
func oddIndex() {
	arr := [5]int{78, 52, 8, 69, 45}
	fmt.Print("Array -", arr)
	fmt.Println("\nprint the odd number ")

	for i := 0; i < 5; i++ {
		if arr[i]%2 != 0 {
			fmt.Print(" ", arr[i])
		}
	}

}

// Print the negative number in the Array
func negative() {
	arr := [5]int{-7, 8, 0, -6, 5}
	fmt.Println("Negative numbers-")
	for i := 0; i < 5; i++ {
		if arr[i] < 0 {
			fmt.Print(" ", arr[i])
		}
	}
}
func main() {
	//largeSmall()
	//evenIndex()
	//oddIndex()
	//negative()
}
