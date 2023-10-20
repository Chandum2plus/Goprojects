package main

import "fmt"

// Implement the program value type.
/*
In Go language, an array is of value type not of reference type. So when the array is assigned to a new variable,
then the changes made in the new variable do not affect the original array
*/

// main function
func main() {
	// Creating an array,which size is Represented by the ellipsis
	ar := [...]int{100, 200, 300, 400, 500}
	fmt.Print("Original Array (BEFORE)-: ", ar)

	// Creating a new variable and initialize with the ar
	arr := ar
	fmt.Println("\nNew array (BEFORE) -: ", arr)

	// change the value at the index 0 to 5
	arr[0] = 500
	fmt.Println("New array (AFTER) -: ", arr)

	// print the original array after the changing
	fmt.Println("Original array (AFTER) -: ", ar)
}
