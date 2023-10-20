package main

import "fmt"

// How ot Access the struct's field
type Car struct {
	name, color, model string
	weightInKg         float64
}

/*
Pointers to a struct
Pointers in Go programming language or Golang is a variable which is used to store the memory address of another variable.
You can also create a pointer to a struct as shown in the below example:
*/
type Employee struct {
	firstName, lastName string
	age, salary         int
}

func main() {
	/*
		c := Car{name: "Rolls Roy's", color: "Orange", model: "Tap maadal", weightInKg: 3000}
		// Accessing the struct fields using dot (.) operator
		fmt.Println("Car Name - ", c.name)
		fmt.Println("Car Color - ", c.color)
		fmt.Println("Car Model - ", c.model)
		fmt.Println("Car Weight -", c.weightInKg)

		// Assigning a new value to a struct field
		c.color = "white"
		// Display the Result
		fmt.Println("Car's Field -", c)
	*/
	// How to Access the pointer structure
	emp := &Employee{
		firstName: "Chandu",
		lastName:  "Kumar",
		age:       22,
		salary:    60000,
	}
	// This is the first way to access the field of the struct

	//fmt.Println("======= FIRST WAY =======")
	//fmt.Println("Name - ", (*emp).firstName, (*emp).lastName)
	//fmt.Println("Age -", (*emp).age)
	//fmt.Println("Salary -", (*emp).salary)

	fmt.Println("======== SECOND WAY ==========") // This is the second way of the access the field of the structure
	fmt.Println("Name -", emp.firstName, emp.lastName)
	fmt.Println("Age -", emp.age)
	fmt.Println("Salary -", emp.salary)
}
