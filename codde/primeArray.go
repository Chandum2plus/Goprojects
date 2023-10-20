package main

import "fmt"

// How to find the prime number from an Array
func primeArray() {
	arr := [...]int{12, 2, 3, 4, 5, 6, 7, 89, 9}

	for i := 0; i < len(arr); i++ {
		flag := 0
		for j := 2; i < len(arr)/2; j++ {
			if arr[i]%j == 0 {
				flag = 1
				break
			}
		}
		if flag == 0 && arr[i] > 1 {
			fmt.Println(arr[i], "Prime")
		}
	}
}
func main() {
	primeArray()
}
