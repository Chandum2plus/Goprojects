package main

import "fmt"

// Main function
func main() {
	// Here, we are going to declare the multidimensional array
	// here i have delcared the array using shorthand declaration
	var arr = [3][3]int{
		{10, 20, 30},
		{90, 80, 70},
		{60, 50, 40},
	}
	// Displaying the Array
	fmt.Println("==== Print the Array =====")
	sum := 0
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			fmt.Print(" ", arr[i][j])
			sum += arr[i][j]
		}
		fmt.Println()
	}
	fmt.Println("Sum of the Array -", sum)
	a := [10]int{} //this is an empty array it has no any element instead of the index.
	// so therefor it will  print the only 10 time zero because its already defined length is 10
	fmt.Println(a)
}
