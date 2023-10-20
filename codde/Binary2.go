package main

import "fmt"

// How to find the element using the binary search
func bin(arr [10]int, key int) int {
	s := 0
	e := len(arr) - 1
	mid := s + (e-s)/2
	for s <= e {
		if arr[mid] == key {
			fmt.Printf("Item found index %d is  = %d", mid, key)
			return mid
		}
		if arr[mid] < key {
			s = mid + 1
		} else {
			e = mid - 1
		}
		mid = s + (e-s)/2

	}
	return -1
	//fmt.Printf("item not found Index %d  is = %d", mid, key)
}
func main() {
	arr := [10]int{12, 23, 34, 45, 56, 78, 89, 90, 100, 102}
	var key int
	fmt.Print("enter the key - ")
	fmt.Scan(&key)
	var value = bin(arr, key)
	//fmt.Printf("item found is =  %d", value)
	fmt.Println("item not found - ", value)
}
