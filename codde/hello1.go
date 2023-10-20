package main

import (
	"fmt"
)

// Delete the element from an array
func delete() {
	arr := [5]int{10, 20, 30, 40, 50}
	i := 3
	// check the index is within array bound
	if i < 0 || i >= len(arr) {
		fmt.Println("item is out of the bound")

	} else { // delete the element
		newlength := 0
		for index := range arr {
			if i != index {
				arr[newlength] = arr[index]
				newlength++
			}

		}
		newArray := arr[:newlength]
		fmt.Println("after deleted - ", newArray)
	}

}

// sum of the total number without using the Array
func sumOfItem() {
	var num1 int
	fmt.Print("enter the three number - ") // here, taking the input
	fmt.Scan(&num1)
	var sum int // this is for sum of the remainder value
	//var r int   // this is for storing the remainder value
	mod := 0
	for num1 > 0 {
		mod = num1 % 10  // here, it is store the remainder value after modulus of the 10
		sum += mod       // it is  sum of the remainder value
		num1 = num1 / 10 // it will store the number after the dividing by 2
	}
	// here check the condition sum === remainder
	for i := 0; i <= num1; i++ {

	}
	fmt.Println("each digit of sum  =", sum)

}

func main() {
	//	delete()
	sumOfItem()
}
