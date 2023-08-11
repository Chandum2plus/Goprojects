package main

import "fmt"

//	func fibo(num int) int {
//		a := 0
//		b := 1
//		for i := 0; i < num; i++ {
//			temp := a
//			a = b
//			b = temp + a
//		}
//		return a
//	}
//
//	func main() {
//		var num int
//		var result int
//		fmt.Println("Enter the number -")
//		fmt.Scanln()
//		for i := 1; i < num; i++ {
//			result = fibo(num)
//		}
//		fmt.Println("Result -", num, result)
//	}
func fibon(num int) {
	var a, b int
	for a, b = 0, 1; a <= num; a, b = b, a+b {
		fmt.Printf(" %d,", a)
	}
	fmt.Println()
	fmt.Println("Result = ", a)
}
func main() {
	var num int
	fmt.Print("Enter the Number - ")
	fmt.Scan(&num)
	fibon(num)
}
