package main

import "fmt"

// How to merge the two array

func mergeTwoArray() {
	// this is the first Array
	arr := []int{10, 20, 30, 40, 50}
	i := 0
	for i = 0; i < len(arr); i++ {
		//fmt.Println(arr[i])
	}
	// This is the second Array
	arr2 := []int{100, 102, 103, 104, 105}
	j := 0
	for j = 0; j < len(arr2); j++ {
		//fmt.Println(arr2[j])
	}

	var arr3 = make([]int, 10)
	arr3 = append(arr, arr2...)
	fmt.Println("Concatenated -", arr3)
}
func main() {
	mergeTwoArray()
}
