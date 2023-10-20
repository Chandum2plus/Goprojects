package main

import (
	"fmt"
)

// copying  the array using the loop
func loop() {
	originalArray := [5]int{10, 20, 30, 40, 50}
	copyArray := make([]int, len(originalArray))
	for i, value := range originalArray {
		copyArray[i] = value
	}
	fmt.Println("Original Array -", originalArray)
	fmt.Println("Copied Array -", copyArray)
}

// copying the array using the copy function
func copys() {
	originalArray := []int{10, 20, 30, 40, 50}
	copyArray := make([]int, len(originalArray))
	copy(copyArray, originalArray)
	fmt.Println("Original Array -", originalArray)
	fmt.Println("Copied Array -", copyArray)
}

// copying the array using the append function
func appends() {
	originalArray := []int{20, 304, 598, 372, 84}
	copyArray := make([]int, len(originalArray))
	copyArray = append(copyArray, originalArray...)

	fmt.Println("Original Array -", originalArray)
	fmt.Println("Copied Array -", copyArray)
}
func main() {
	//loop()
	//copys()
	appends()

}
