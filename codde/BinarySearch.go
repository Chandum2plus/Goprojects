package main

import (
	"fmt"
)

// find the element of the Array using the binary search
func binary() {
	var key int
	fmt.Print("enter the key - ")
	fmt.Scan(&key)
	arr := [10]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	mid := 0
	first := 0
	last := 9
	flag := 1
	mid = (first + last) / 2
	for i := 0; i < len(arr); i++ {

	}
	for first = 0; first <= last; {
		if arr[mid] < key {
			first = last + 1
		} else if arr[mid] == key {
			flag = 1
			fmt.Print("item found -", key)
			break
		} else {
			last = mid - 1
		}
		mid = (first + last) / 2
	}
	if flag == 0 {
		fmt.Println("Item not found")
	}
}

func main() {
	binary()
}
