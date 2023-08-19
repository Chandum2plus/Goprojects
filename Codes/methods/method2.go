package main

import "fmt"

// Declaring the structures
type student struct {
	name   string
	branch string
}
type teacher struct {
	language string
	marks    int
}

func (s student) show() {
	fmt.Println("Name of the Student - ", s.name)
	fmt.Println("Name of the Branch - ", s.branch)
}
func (t teacher) show() {
	fmt.Println("Name of the Language - ", t.language)
	fmt.Println("Student's Marks - ", t.marks)
}
func main() {
	val1 := teacher{"Golang", 69}
	val2 := student{"Chandu kumar", "Compute Application"}
	val1.show()
	val2.show()

}
