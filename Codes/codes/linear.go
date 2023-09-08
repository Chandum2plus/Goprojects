package main

import (
	"fmt"
)

func liniar() {
	var key int
	fmt.Print("enter the key - ")
	fmt.Scan(&key)
	flag := 0
	arr := [5]int{12, 34, 56, 87, 88}
	for i := 0; i < len(arr); i++ {
		if arr[i] == key {
			flag = 1
			break
		}
	}
	if flag == 1 {
		fmt.Println("Key found -", key)
	} else {
		fmt.Println("bhosdi ke  aukad me rah ke search karo")
	}
}
func main() {
	liniar()

}
