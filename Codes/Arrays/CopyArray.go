package main

import "fmt"

// How, to copy an array into the another array
func main() {
	// Here, we are initializing the array with the help of the shortHand operator
	arr1 := [5]string{"Chandu", "sonal", "saurabh", "somi", "monika"}
	// copying the array into the new variable
	// Here, element are the passed by the value
	CpAr := arr1
	fmt.Println("Original Array - ", arr1)
	fmt.Println("Copied Array -", CpAr)
	//CpAr[0] = "Anjali" // You can change made in the  copied array, and without affect the original array
	//	fmt.Println("Copied Changed or not - ", CpAr)

	// Here, when we copy an array
	// into another array by value
	// then the changes made in original
	// array do not reflect in the
	// copy of that array

	arr1[0] = "Anjali"
	fmt.Println("Original Array After Copy -", arr1)
	fmt.Println("Copied Array After -", CpAr)
}
