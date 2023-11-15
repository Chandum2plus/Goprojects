package main

import "fmt"

// example of the interface using the structure

// this is the interface named as arith.
type arith interface {

	// these are the methods
	sum() float64
	mul() float64
}
type param struct {
	num1 float64
	num2 float64
}

func (p param) sum() float64 {
	return p.num1 + p.num2
}

func (p param) mul() float64 {
	return p.num2 * p.num2
}

func Display(a arith) {
	fmt.Println("sum of two number -", a.sum())
	fmt.Println("Multiplication of the two number -", a.mul())
}
func main() {
	//s := param{
	//	num1: 34,
	//	num2: 45,
	//}
	var s param
	fmt.Print("Enter the first number -")
	fmt.Scan(&s.num1)
	fmt.Print("Enter the second number -")
	fmt.Scan(&s.num2)

	Display(s)

}
