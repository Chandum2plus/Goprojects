package main

import "fmt"

// find the  Greatest and smallest number an Array
func greatest() {
	// Creating the Array
	arr := [5]int{2, 23, 45, 67, 7}
	fmt.Println("Original Array - ", arr)
	gre := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > gre {
			gre = arr[i]
		}
	}
	sma := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] < sma {
			sma = arr[i]
		}
	}

	fmt.Println("Greatest - ", gre)
	fmt.Println("Smallest - ", sma)
}
func main() {
	greatest()

}
