package main

import (
	"fmt"
)

// initializing the nested structure
//
//	type Details struct {
//		details1 Teacher
//	}
func show() {
	// initializing the structure
	type Student struct {
		name    string
		address string
		class   string
		roll    int
	}
	type Teacher struct {
		faculty        string
		qualification  string
		courseName     string
		mob            int
		detailsStudent Student
	}
	res := Teacher{

		//details:  Student{name: "Chandu", address: "Adma", class: "BCA", roll: 68},
		//details1: Teacher{faculty: "Fundamental", qualification: "Post Graduate", courseName: "MCA", mob: 1234567890},,
		detailsStudent: Student{},
	}
	fmt.Print("Enter Name of Student - ")
	fmt.Scan(&res.detailsStudent.name)
	fmt.Print("Enter the Address - ")
	fmt.Scan(&res.detailsStudent.address)
	fmt.Print("Enter the Class - ")
	fmt.Scan(&res.detailsStudent.class)
	fmt.Print("Enter the Roll - ")
	fmt.Scan(&res.detailsStudent.roll)

	fmt.Print("Enter the Faculty's Name - ")
	fmt.Scan(&res.faculty)
	fmt.Print("Enter the Qualification - ")
	fmt.Scan(&res.qualification)
	fmt.Print("Enter the Course Name - ")
	fmt.Scan(&res.courseName)
	fmt.Print("Enter the Mobile - ")
	fmt.Scan(&res.mob)

	fmt.Println("\n====== RESULT =======")
	fmt.Println("Name -", res.detailsStudent.name)
	fmt.Println("Address -", res.detailsStudent.address)
	fmt.Println("Class - ", res.detailsStudent.class)
	fmt.Println("Roll - ", res.detailsStudent.roll)
	fmt.Println("\n===== TEACHER'S DETAILS =========")
	fmt.Println("Faculty -", res.faculty)
	fmt.Println("Qualification - ", res.qualification)
	fmt.Println("Course - ", res.courseName)
	fmt.Println("Mobile  -", res.mob)
}
func show1() {
	type student struct {
		name   string
		branch string
		year   int
	}
	type teacher struct {
		name    string
		subject string
		exp     int
		detail  student
	}
	res := teacher{
		name:    "Subodh kumar",
		subject: "Go language",
		exp:     10,
		detail:  student{"Chandu", "ECS", 3},
	}
	fmt.Println("\n ======Details of the Teacher ======= ")
	fmt.Println("Name of the Teacher -", res.name)
	fmt.Println("Subject -", res.subject)
	fmt.Println("Experience -", res.exp)
	fmt.Println("\n=========== Details of the Student =============")
	fmt.Println("Name of the Student -", res.detail.name)
	fmt.Println("Name of the Branch -", res.detail.branch)
	fmt.Println("Year -", res.detail.year)
}
func main() {
	//show()
	show1()
}
