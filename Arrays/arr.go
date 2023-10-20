package main

import "fmt"

func main() {
	arr := [5]int{12, 34, 56, 43, 32}
	sum := 0
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i])
		sum = arr[i]
	}
	fmt.Printf("\nSum = %d", sum)
}
