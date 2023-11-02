package main

import (
	"fmt"
)

/*
12. Write a C program to read the roll no, name and marks of three subjects and calculate the total, percentage and division.
Test Data :
Input the Roll Number of the student :784
Input the Name of the Student :James
Input the marks of Physics, Chemistry and Computer Application : 70 80 90
Expected Output :
Roll No : 784
Name of Student : James
Marks in Physics : 70
Marks in Chemistry : 80
Marks in Computer Application : 90
Total Marks = 240
Percentage = 80.00
Division = First
*/
type student struct {
	name         string
	roll         int
	phy          float64
	chem         float64
	computerApps float64
}

func percentage() {
	st := student{}
	fmt.Print("Enter the Name - ")
	fmt.Scan(&st.name)
	fmt.Print("Enter the Roll - ")
	fmt.Scan(&st.roll)
	fmt.Print("Enter the Physic's Marks - ")
	fmt.Scan(&st.phy)
	fmt.Print("Enter the Chemistry's Marks - ")
	fmt.Scan(&st.chem)
	fmt.Print("Enter the Computer's Marks - ")
	fmt.Scan(&st.computerApps)
	var total float64
	var perc float64
	total = st.phy + st.chem + st.computerApps
	perc = total / 300 * 100
	if perc >= 60 {
		fmt.Println("first Division", perc)
	} else if perc >= 48 && perc < 60 {
		fmt.Println("Second Division", perc)
	} else if perc < 48 && perc >= 30 {
		fmt.Println("Third Division", perc)
	} else {
		fmt.Println("Fail", perc)
	}
}
func main() {
	percentage()
}
