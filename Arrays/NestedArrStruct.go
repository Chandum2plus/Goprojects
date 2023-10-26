package main

import "fmt"

// find the prime number
func prinme() {
	var num int
	flag := 0
	fmt.Print("Enter the number - ")
	fmt.Scan(&num)
	for i := 0; i < num/2; i++ {
		if num%2 == 0 {
			flag++
			break
		}
	}
	if flag == 0 && num != 1 {
		fmt.Println("Prime number =", num)
	} else {
		fmt.Println("Not prime number =", num)
	}
}
func main() {
	prinme()
}
