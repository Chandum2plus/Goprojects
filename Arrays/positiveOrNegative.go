package main

import "fmt"

// find the postive or negative number in the array
func posNeg() {

	// this is the fisrt Approach
	var arr = [5]int{0, 2, 3, 4, 0}
	for i := 0; i < len(arr); i++ {
		if arr[i] > 0 {
			fmt.Println("Positve Number -", arr[i])
		} else {
			fmt.Println("Negative Number -", arr[i])
		}
	}
}

// This is the second approach
func posNeg2() {

	counterP := 0
	counterN := 0
	var size int
	fmt.Print("Enter the size of the Array - ")
	fmt.Scan(size)
	var arr = make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Print("Enter Positive Negative Number - ")

		fmt.Scan(&arr[i])
	}
	for i := 0; i < size; i++ {
		if arr[i] > 0 {
			counterP++
		} else {
			counterN++
		}
	}
	fmt.Println("Total Positive Number - ", counterP)
	fmt.Println("Total Negative Number - ", counterN)
}
func main() {
	//posNeg()
	posNeg2()
}
