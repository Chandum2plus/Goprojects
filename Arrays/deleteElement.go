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
	var num1, num2, num3 int
	fmt.Print("enter the three number - ")
	fmt.Scan(&num1, &num2, &num3)
	sum := 0
	sum = (num1 + num2 + num3)

	fmt.Println("sum  =", sum)

}
func main() {
	//	delete()
	sumOfItem()
}
