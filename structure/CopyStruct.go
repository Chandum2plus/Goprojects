package main

import "fmt"

// How to copy the structure type using the value and pointer reference
type cop struct {
	length  int
	breadth int
	color   string
}

func copi() {
	c1 := cop{
		length:  12,
		breadth: 23,
		color:   "Red",
	}
	fmt.Println("C1 = ", c1)
	c2 := c1
	c2.color = "Green"
	fmt.Println("C2 = ", c2)

	c3 := &c1
	c3.length = 39
	c3.breadth = 57
	c3.color = "Yellow"
	fmt.Println("C3 = ", c3)
	fmt.Println("C1 = ", c1)
}
func main() {
	copi()

}
