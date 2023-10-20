package main

import "fmt"

// Example of the method can accept both pointer and value

type developers struct { // Declared the structure
	name    string
	company string
	job     string
	salary  float64
}

// declaring the method
func (d *developers) devlop(acompany string) {
	(*d).company = acompany
}
func (d developers) devlop2() {
	d.name = "Chandu"
	d.company = "tcs"
	d.job = "frontend"
	d.salary = 34999.56
	fmt.Println("Name - ", d.name)
	fmt.Println("Company - ", d.company)
	fmt.Println("Job - ", d.job)
	fmt.Println("Salary - ", d.salary)
}
func main() {
	res := developers{
		name:    "Chandu_kumar",
		company: "infosys",
		job:     "frontend",
		salary:  23344.90,
	} //
	fmt.Println("===Developers details before ===")
	fmt.Println("Name - ", res.name)
	fmt.Println("Company - ", res.company)
	fmt.Println("Job - ", res.job)
	fmt.Println("Salary - ", res.salary)

	// calling the develop method pointer with value
	res.devlop("cisco")
	fmt.Println("company Name (After) - ", res.company)
	// calling the develop2 method pointer with the value
	(&res).devlop2()
	fmt.Println("Name of the developer (After) - ", res.name)
}
