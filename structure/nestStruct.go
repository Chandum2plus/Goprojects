package main

import "fmt"

/*
Nested Struct Type
Struct can be nested by creating a Struct type using other Struct types as the type for the fields of Struct.
Nesting one struct within another can be a useful way to model more complex structures.
*/
type Salary struct {
	HRA, BASIC, TA float64
}
type Employ struct {
	Name          string
	Age           float64
	Email         string
	MonthlySalary []Salary
}

func nestedStruct() {
	e := Employ{
		Name:  "Chandu_kumar",
		Age:   23,
		Email: "chandukumarbca@gmail.com",
		MonthlySalary: []Salary{
			{
				HRA:   5000,
				BASIC: 50000,
				TA:    2000,
			},
			{
				HRA:   6000,
				BASIC: 55000,
				TA:    2100,
			},
			{
				HRA:   7000,
				BASIC: 60000,
				TA:    2200,
			},
		},
	}
	fmt.Println("Name - ", e.Name)
	fmt.Println("Age - ", e.Age)
	fmt.Println("Email - ", e.Email)
	fmt.Println("Salary1 = ", e.MonthlySalary[0])
	fmt.Println("Salary2 = ", e.MonthlySalary[1])
	fmt.Println("Salary3 = ", e.MonthlySalary[2])

}
func main() {
	nestedStruct()
}
