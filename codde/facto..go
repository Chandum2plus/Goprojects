package main

import "fmt"

// find the factorial of number...../
func facto() {
	var n int
	fmt.Print("Enter the number - ")
	fmt.Scan(&n)
	fact := 1
	for i := 1; i <= n; i++ {
		fact = fact * i
	}
	fmt.Println("Facto. =", fact)
}
func main() {
	facto()
}
