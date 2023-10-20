package main

import "fmt"

// embed interface
type Geometory interface {
	Edges() int
}
type Polygons interface {
	Geometory
}
type pentagons int
type hexagons int
type octagons int
type decagons int

func (p pentagons) Edges() int {
	return 5
}
func (h hexagons) Edges() int {
	return 6
}
func (o octagons) Edges() int {
	return 8
}
func (d decagons) Edges() int {
	return 10
}
func main() {
	p := new(pentagons)
	h := new(hexagons)
	o := new(octagons)
	d := new(decagons)

	var poly = [...]Polygons{p, h, o, d}
	for i := range poly {
		fmt.Println(poly[i].Edges())
	}

}
