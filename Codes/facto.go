package main

import "fmt"

func main() {
	var num int
	fact := 1
	fmt.Print("Enter the number -")
	fmt.Scan(&num)
	for i := 1; i <= num; i++ {
		fact = fact * i

	}
	fmt.Println("factorial = ", fact)
}
