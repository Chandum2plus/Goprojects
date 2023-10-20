package main

import "fmt"

// How to compare the two data in a structure
type comp struct {
	length  int
	Breadth int
	color   string
}

func compare() {
	c := comp{
		length:  12,
		Breadth: 20,
		color:   "Red",
	}
	c1 := comp{12, 20, "Red"}
	if c == c1 {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
	c2 := new(comp)
	var c3 = &comp{}
	if c2 == c3 {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}
func main() {
	compare()
}
