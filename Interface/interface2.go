package main

import "fmt"

// implementing an interface

// Creating an interface

type Tank interface {
	Tarea() float64
	Volume() float64
}

// Creating the struct for variable

type Myvalue struct {
	radius float64
	height float64
}

// implementing the method

func (m Myvalue) Tarea() float64 {
	return 2*m.radius*m.height + 2*3.14*m.radius*m.radius
}
func (m Myvalue) Volume() float64 {
	return 3.14 * m.radius * m.radius * m.height
}

func main() {
	var t Tank
	t = Myvalue{10, 14}
	fmt.Println("Area of tank - ", t.Tarea())
	fmt.Println("Area of volume - ", t.Volume())
}
