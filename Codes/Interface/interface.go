package main

import (
	"fmt"
)

// Here, declaring an interface of Employee type
type Employee interface {
	PrintName(name, compName string)
	PrintSalary(basic int, tax int) int
	PrintAge(age float64, house int)
}

// this is the user defined type named as Emp
type Emp int

// PrintName method to print the emp. name and comp Name
func (e Emp) PrintName(name, compName string) {
	fmt.Println("Employee Id -", e)          // this is for id
	fmt.Println("Employee Name - ", name)    // this is for name
	fmt.Println("Company Name - ", compName) // this is for company name
}

// Printsalary method to print the calculated salary
func (e Emp) PrintSalary(basic int, tax int) int {
	var salary = (basic * tax) / 100
	return basic - salary
}

// PrintAge method to print the age and house number of emp
func (e Emp) PrintAge(age float64, house int) {
	fmt.Println("Employee Age -", age) // this is for Age
	fmt.Println("House No. -", house)  // this is for house
}

// This is the Main function
func main() {
	var e1 Employee                                             // it is the variable of Employee type
	e1 = Emp(1)                                                 // passing the id No
	e1.PrintName("Chandu kumar", "TCS")                         // passing the name and the comp Name
	e1.PrintAge(22.8, 35)                                       // passing the age and house number
	fmt.Println("Employee Salary -", e1.PrintSalary(255000, 5)) // Here passing and printing the salary included tax paid

}
