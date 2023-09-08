package main

import "fmt"

// sorting the element of an array in ascending order
func sorting() {

	// Declaring the Array
	arr := [5]int{12, 31, 34, 5, 3}
	temp := 0
	fmt.Print("Unsorted Array - ", arr)
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[i] < arr[j] {
				temp = arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}

		}
	}
	fmt.Print("\nSorted Array - ")
	for i := 0; i < len(arr); i++ {
		fmt.Print(" ", arr[i])
	}
}

// sorting the element of an array in descending order
func sortingDec() {

	// declaring an array
	arr := [5]int{1, 2, 9, 4, 6}
	fmt.Print("Unsorted - ", arr)
	temp := 0
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[j] < arr[i] {
				temp = arr[j]
				arr[j] = arr[i]
				arr[i] = temp
			}
		}
	}
	fmt.Print("\nSorted -")
	for i := 0; i < len(arr); i++ {
		fmt.Print(" ", arr[i])
	}
}
func main() {
	//sorting()
	sortingDec()
}
