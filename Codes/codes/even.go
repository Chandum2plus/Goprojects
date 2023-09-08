package main

import (
	"fmt"
)

// find the even number of  the Array
func evens() {
	arr := [5]int{12, 4, 56, 23, 43}

	ev := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 == 0 {
			ev = arr[i]
			fmt.Println(" ", ev)

		}
	}
	//fmt.Println("even - ", ev)
}
func main() {
	evens()
}
