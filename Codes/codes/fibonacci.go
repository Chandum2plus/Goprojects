package main

import "fmt"

func fibo() {
	var num, c int
	a := 0
	b := 1
	fmt.Print("Enter the number - ")
	fmt.Scan(&num)
	fmt.Print("Fibonacci - ")
	for i := 0; i <= num; i++ {
		fmt.Print(" ", a)
		c = a + b
		a = b
		b = c
	}
	fmt.Println("\nseries = ", c)
}
func fiib() {
	var num int
	fmt.Print("enter the number - ")
	fmt.Scan(&num)
	a := 0
	b := 1
	c := 0
	for i := 0; i <= num; i++ {
		fmt.Print(" ", a)
		c = a + b
		a = b
		b = c

	}
	fmt.Println("\nfibbo", c)
}
func main() {
	//fibo()
	fiib()
}
