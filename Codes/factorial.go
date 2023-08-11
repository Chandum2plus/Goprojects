package main

import "fmt"

func facts() {
	var num int
	fmt.Print("Enter the Number - ")
	fmt.Scan(&num)
	fact := num * num * num
	fmt.Println("Factorial of the Number - ", fact)
}
func main() {
	facts()
}
