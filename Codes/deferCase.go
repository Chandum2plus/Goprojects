package main

import (
	"fmt"
)

func mul(a1, a2 int) int {
	res := a1 * a2
	fmt.Println("Result -", res)
	return 0
}
func show() {
	fmt.Println("Hello Chandu")
}
func add(num1, num2 int) int {
	res := num1 + num2
	fmt.Println("Result- ", res)
	return 0
}
func main() {
	//mul(23, 45)
	//defer mul(23, 48)
	//defer mul(56, 78)
	//show()
	fmt.Println("Start")
	defer fmt.Println("End")
	defer add(23, 45) // output = 68
	defer add(89, 45) // output = 134
	defer add(90, 45) // output = 135

}
