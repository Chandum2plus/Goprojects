package main

import "fmt"

// <============ Polymorphism=============>
type Geometry interface {
	Edge() int
}
type pentagon struct{} // pentagon define a Geometry object
type hexagon struct{}  // define a Geometry object
type octagon struct{}  // define a Geometry object
type decagon struct{}  // define a Geometry object

// Edge implement the Geometry interface
func (p pentagon) Edge() int {
	return 5
}
func (h hexagon) Edge() int {
	return 6
}
func (o octagon) Edge() int {
	return 8
}
func (d decagon) Edge() int {
	return 10
}

// calculate the parameter of the object
func parameter(g Geometry, value int) int {

	num := g.Edge()
	calculation := num * value
	return calculation
}
func main() {

	p := new(pentagon)
	h := new(hexagon)
	o := new(octagon)
	d := new(decagon)

	g := [...]Geometry{p, h, o, d}
	for _, i := range g {
		fmt.Println(parameter(i, 5))
	}

}
