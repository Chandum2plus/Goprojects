package main

import "fmt"

// Implement the details of the student

// creating an interface
type Student interface {
	classDetails(name, course string)
	Address(vill, ps, po, dist string)
}

// creating the structure
type Details int

// implementing the class() method
func (d Details) classDetails(name, course string) {
	fmt.Println("ID - ", d)
	fmt.Println("Name - ", name)
	fmt.Println("course - ", course)

}
func (d Details) Address(vill, ps, po, dist string) {
	fmt.Println("Village - ", vill)
	fmt.Println("Police station - ", ps)
	fmt.Println("Post office - ", po)
	fmt.Println("District - ", dist)
}

// main function
func main() {
	var st Student
	st = Details(1)
	st.classDetails("Chandu", "BCA")
	st.Address("Adma", "Obra", "Krashawa", "Aurangabad")

}
