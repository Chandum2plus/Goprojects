// Implementing the interface
/*
package main

import "fmt"

// create an interface
type Employ interface {
	PrintName(name string) // method with string return type
	//	PrintAdd(id int)   // method with parameter
	PrintSalary(basic, tax int) int // method with parameter and return type
}
type E int

// PrintName() method
func (e E) PrintName(name string) {
	fmt.Println("Employee Id :-", e)
	fmt.Println("Employee Name :-", name)

}

// PrintSalary() method
func (e E) PrintSalary(basic int, tax int) int {
	var salry = (basic * tax) / 100
	return basic - salry

}
func main() {
	var e1 Employ
	e1 = E(1)
	e1.PrintName("Chandu kumar")
	fmt.Println("Employee Salary :- ", e1.PrintSalary(25000, 5))

}

*/

// implementing the interface
package main

import "fmt"

// Create an interface
type Employeee interface {
	PrintName(name string)
	PrintSalary(basic, tax int) int
}
type Emps int

// printName() method
func (e Emps) PrintName(name string) {
	fmt.Println("Employee Id :-", e)
	fmt.Println("Employee Name :-", name)
}

// calculate the employee salary PrintSalary() method
func (e Emps) PrintSalary(basic int, tax int) int {
	var salary = (basic * tax) / 100
	return basic - salary
}
func main() {
	var e1 Employeee                                            //this is the instance variable of the Employeee
	e1 = Emps(2)                                                // it store the id of employee
	e1.PrintName("Chandu kumar")                                // calling the PrintName() function passing the parameter
	fmt.Println("Employee salary :-", e1.PrintSalary(28000, 5)) // calling the PrintSalary() method passing the parameter

}
