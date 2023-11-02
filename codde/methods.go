package main

import "fmt"

// Example of the method using the return type

// this the method named as method which has two parameter length and width which type is float data-type,
// and it also Returns the value
func method(length, width float64) float64 {

	Ar := length * width // Here, calculating the result
	return Ar            // Here, returning the value
}

// This is the main function
func main() {
	// taking the two variable named as a b which type is float
	var a, b float64

	// Here, taking the user input
	fmt.Print("Enter the length - ")
	fmt.Scan(&a)
	fmt.Print("Enter the width - ")
	fmt.Scan(&b)

	// Here, calling the function and it also displays the Result
	fmt.Println("Area=", method(a, b))
}
