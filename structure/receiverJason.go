package main

import "fmt"

type Emplo struct {
	Name           string
	Age            int
	Email          string
	MonthalySalary []Salarry
}
type Salarry struct {
	BASIC, HRA, TA float64
}

func (e Emplo) getinfo() string {
	fmt.Println("Name -", e.Name)
	fmt.Println("Age -", e.Age)
	fmt.Println("Email -", e.Email)

	for _, i := range e.MonthalySalary {
		fmt.Println("=======================")
		fmt.Println("Basic =", i.BASIC)
		fmt.Println("HRA =", i.HRA)
		fmt.Println("TA =", i.TA)
	}
	return "...................."
}
func main() {

	emp := Emplo{
		Name:  "Chandu_kumar",
		Age:   23,
		Email: "chandukumarbca@gmail.com",
		MonthalySalary: []Salarry{
			{
				BASIC: 30000,
				HRA:   5000,
				TA:    2000,
			},
			{
				BASIC: 40000,
				HRA:   6000,
				TA:    2100,
			},
			{
				BASIC: 50000,
				HRA:   7000,
				TA:    2200,
			},
		},
	}
	fmt.Println(emp.getinfo())
}
