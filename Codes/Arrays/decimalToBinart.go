package main

import (
	"fmt"
)

// convert decimal to binary and store into the Array
func dec() {
	var num int
	var binarr []int // this is the Array variable which store the binary number

	// Taking the user input a number
	fmt.Print("Enter the Decimal Number - ")
	fmt.Scan(&num)

	for num != 0 { // yeh tb tk chalega jab tk zero n mil jaye
		binarr = append(binarr, num%2)
		num = num / 2
	}
	if len(binarr) == 0 {
		fmt.Printf("%d", 0)
	} else {
		for i := len(binarr) - 1; i >= 0; i-- { // Here, Reversing the Array
			fmt.Printf("%d", binarr[i])
		}
		fmt.Println()
	}

}
func main() {
	dec()
}
