package main

import "fmt"

// How to an array into the another by reference
func main() {
	// Here, we are initializing the array using the shortHand Declaration
	ar := [5]int{10, 20, 30, 40, 50}

	// Copying the array into the new variable
	// Here, all elements passed by the reference
	cpAr := &ar
	fmt.Println("Original Array - ", ar)
	fmt.Println("Copied Array -", *cpAr)

	// Here, when we copy an array
	// into another array by reference
	// then the changes made in original
	// array will reflect in the
	// copy of that array
	ar[4] = 10000
	fmt.Println("Original Array After Change made -", ar)
	fmt.Println("Copied Array After Change made -", *cpAr)

}
