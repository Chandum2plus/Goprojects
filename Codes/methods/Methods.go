package main

import "fmt"

// method with the struct type
type developer struct {
	name   string
	branch string
	job    string
	salary int
}

// Method with a Receiver author type
func (d developer) show() { // this is the receiver method
	fmt.Println("\n======== RESULT =========")
	fmt.Println("Developer Name - ", d.name)
	fmt.Println("Developer Branch - ", d.branch)
	fmt.Println("Developer Job - ", d.job)
	fmt.Println("Developer Salary - ", d.salary)
}

// defining a method with the non struct type receiver
type data float64 // Here we are created the new datatype as like which named is data

func (d1 data) multiply(d2 data) data {
	return d1 * d2

}

// Method with the pointer receiver
type devloper struct {
	name    string
	company string
	job     string
	salary  float64
}

func (p *devloper) show(acompany string) {
	(*p).company = acompany
}
func main() {

	/* // This is the first way of the sending the data to receiver method
	res := developer{
		name:   "Chandu kumar",
		branch: "CISCO",
		job:    "Backend Developer",
		salary: 600000,
	}

	*/
	/*
		// this is the second or user input way of the sending the data to receiver method ....
		r := developer{} // from here, i am sending the data .......
		fmt.Print("Enter the Developer Name - ")
		fmt.Scan(&r.name)
		fmt.Print("Enter the Developer Branch - ")
		fmt.Scan(&r.branch)
		fmt.Print("Enter the Developer Job - ")
		fmt.Scan(&r.job)
		fmt.Print("Enter the Developer Salary - ")
		fmt.Scan(&r.salary)
		// Here, calling the method
		r.show()
	*/
	/*
		var a, b data
		fmt.Print("Enter the first number - ")
		fmt.Scan(&a)
		fmt.Print("Enter the second number - ")
		fmt.Scan(&b)
		res := data.multiply(a, b) // ye usi variable ko accept  krta hai jo kisi type se  bna hota hai
		// example:- type data int
		// var a,b data - data is the also a datatype
		fmt.Println("Multiplied -", res)
	*/
	// initializing the value of the devloper structure
	res := devloper{}
	fmt.Print("Enter the Employee Name - ")
	fmt.Scan(&res.name)
	fmt.Print("Enter the Company Name - ")
	fmt.Scan(&res.company)
	fmt.Print("Enter the Job Name - ")
	fmt.Scan(&res.job)
	fmt.Print("Enter the Salary - ")
	fmt.Scan(&res.salary)
	fmt.Println("\n======== RESULT =========")
	fmt.Println("Name -", res.name, "\nCompany - ", res.company, "\nJob - ", res.job, "\nSalary - ", res.salary)

	// Here , i am creating a pointer variable
	p := &res
	// calling the show method
	fmt.Println(" =========== After =============")
	res.show("cisco")
	fmt.Println("Name - ", p.name)
	fmt.Println("Company -", p.company)
	fmt.Println("Job - ", p.job)
	fmt.Println("Salary - ", p.salary)

}
