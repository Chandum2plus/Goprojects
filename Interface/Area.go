package main

import (
	"fmt"
)

// creating an interface
type Shape interface {
	area() float64
}

// creating the structure
type Rectangle struct {
	length, breath float64
}

func (r Rectangle) area() float64 {
	return r.length * r.breath
}
func calculate(s Shape) {
	fmt.Println("Area = ", s.area())
}

func main() {
	// Create the variable of the Rectangle type
	rect := Rectangle{}
	fmt.Print("enter length - ")
	fmt.Scan(&rect.length)
	fmt.Print("enter the breath - ")
	fmt.Scan(&rect.breath)
	// calling the calculate()method with struct variable of the rect
	calculate(rect)
}
