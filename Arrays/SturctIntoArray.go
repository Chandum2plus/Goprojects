package main

import (
	"fmt"
)

// How to store the Struct into the Array in golang
// Static Initialization of a Golang array of structs

func structIntoArray() {

	// Creating a Structure
	type student struct {
		name  string
		roll  int
		class string
	}
	var mst = student{}
	mst.name = "Chintu kumar"
	mst.roll = 12
	mst.class = "MCA"
	fmt.Print("mst =", mst)

	// Here, we are storing the structure into the array
	var st = [3]student{
		{"Chandu", 23, "MCA"},
		{"Saurav ", 45, "MCA"},
		{"Sonal", 33, "BCA"},
	}
	fmt.Println("\nStudent Array =", st) // Printing the Result
}

// Dynamic Initialization of a Golang array of structs
func dynamicStoring() {

	// creating the Structure
	type student struct {
		name  string
		roll  int
		class string
	}

	// storing into the Array
	var stArr = [3]student{}

	// iterating the student using the loop
	for i := 0; i < len(stArr); i++ {
		fmt.Println("Enter the Name - ")
		fmt.Scanln(&stArr[i].name)
		fmt.Println("Enter the Roll - ")
		fmt.Scanln(&stArr[i].roll)
		fmt.Println("Enter the Class - ")
		fmt.Scanln(&stArr[i].class)
	}
	fmt.Println("Result - ", stArr)
}

func main() {
	//structIntoArray()
	dynamicStoring()
}
