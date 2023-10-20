package main

import "fmt"

func serch(arr [10]int, key int) bool {

	//i := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == key {
			return true
		}
	}
	return false

}
func main() {
	//arr := [10]int{10, 20, 32, 43, 56, 76, 44, 33, 22, 12}
	//
	//var key int
	//fmt.Print("enter the key - ")
	//fmt.Scan(&key)
	//found := (serch(arr, 10))
	//if found {
	//	fmt.Println("Item found")
	//} else {
	//	fmt.Println("item not found")
	//}
	linear()
}
func linear() {
	var key int
	fmt.Print("enter the key - ")
	fmt.Scan(&key)
	arr := [10]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	flag := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == key {
			flag = 1
			break
		}
	}
	if flag == 1 {
		fmt.Println("key found")
	} else {
		fmt.Println("key not found")
	}
}
