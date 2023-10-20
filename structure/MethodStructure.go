package main

import "fmt"

// Add Method to Struct Type
// You can also add methods to struct types using a method receiver. A method EmpInfo is added to the Employee struct.
type salari struct {
	Basic, HRA, TA float64
}
type employee struct {
	FirstName     string
	LastName      string
	Email         string
	Age           float64
	monthlysalary []salari
}

func (e employee) empInfo() string {
	fmt.Println("Name - ", e.FirstName, e.LastName)
	fmt.Println("Email - ", e.Email)
	fmt.Println("Age - ", e.Age)
	for _, info := range e.monthlysalary {
		fmt.Println("=============")
		fmt.Println("Basic salary - ", info.Basic)
		fmt.Println("HRA - ", info.HRA)
		fmt.Println("TA - ", info.TA)
	}
	return "......."

}
func main() {
	e := employee{
		FirstName: "Chandu",
		LastName:  "Kumar",
		Email:     "chandukumarbca@gmail.com",
		Age:       22.9,
		monthlysalary: []salari{
			{
				Basic: 25000,
				HRA:   500,
				TA:    200,
			},
			{
				Basic: 26000,
				HRA:   5000,
				TA:    2100,
			},
			{
				Basic: 29000,
				HRA:   5000,
				TA:    2200,
			},
		},
	}
	fmt.Println(e.empInfo()) //function called
}
