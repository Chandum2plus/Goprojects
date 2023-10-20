package main

import "fmt"

// How to access the nested structure
type salary struct {
	Basic, HRA, TA float64
}
type Employee struct {
	firstName, LastName, email string
	age                        float64
	monthlySalary              []salary
}

func nested() {
	e := Employee{
		firstName: "Chandu ",
		LastName:  "Kumar",
		email:     "chandukumarbca@gmail.com",
		age:       22.9,
		monthlySalary: []salary{
			{
				Basic: 25000,
				HRA:   5000,
				TA:    2000,
			},
			{
				Basic: 26000,
				HRA:   5000,
				TA:    2100,
			},
			{
				Basic: 29000,
				HRA:   5000,
				TA:    2300,
			},
		},
	}
	fmt.Println("Name - ", e.firstName, e.LastName)
	fmt.Println("Age - ", e.age)
	fmt.Println("email - ", e.email)
	fmt.Println("First salary - ", e.monthlySalary[0])
	fmt.Println("Second salary - ", e.monthlySalary[1])
	fmt.Println("Third salary - ", e.monthlySalary[2]) // this will print index wise  like this given below
	/*
	   Name -  Chandu  Kumar
	   Age -  22.9
	   email -  chandukumarbca@gmail.com
	   First salary -  {25000 5000 2000}
	   Second salary -  {26000 5000 2100}

	*/
}

func main() {
	nested()

}
