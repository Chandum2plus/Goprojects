package main

import "fmt"

type Polygon interface {
	Perimeter()
}
type Object interface {
	NumberOfSide()
}
type penta int

func (p penta) Perimeter() {
	fmt.Println("Perimeter of the Pentagon - ", 5*p)
}
func (p penta) NumberOfSide() {
	fmt.Println("Pentagon has 5 sides ")
}

func main() {
	var p Polygon = penta(50)
	p.Perimeter()
	var o penta = p.(penta)
	o.NumberOfSide()

	var obj Object = penta(50)
	obj.NumberOfSide()
	var pen penta = obj.(penta)
	pen.Perimeter()
}
