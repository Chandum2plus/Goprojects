package main

import (
	"fmt"
)

// find the prime number using the user input

func primeNumb() {
	var num int
	fmt.Print("Enter the number - ")
	fmt.Scan(&num)
	temp := 0
	flag := 0
	temp = num / 2
	if num == 0 || num == 1 {
		fmt.Printf("%d is Not prime", num)
	} else {
		for i := 2; i <= temp; i++ {
			if num%i == 0 {
				fmt.Printf("%d is Not Prime", num)
				flag = 1
				break
			}
		}
		if flag == 0 {
			fmt.Printf("%d is Prime number", num)
		}
	}

}
func main() {
	primeNumb()
}
