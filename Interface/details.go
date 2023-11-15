package main

import "fmt"

// Example of the interface

// interface named as details
type details interface {
	student() string
}

// structure named as data
type data struct {
	name   string
	id     int
	class  string
	roll   int
	course string
}

func (d data) student() {
	fmt.Println("Enter the Name -", d.name)
	fmt.Println("Enter the Id -", d.id)
	fmt.Println("Enter the Class -", d.class)
	fmt.Println("Enter the Roll -", d.roll)
	fmt.Println("Enter the Course -", d.course)
}

func display(dd details) {

}

func main() {

}
