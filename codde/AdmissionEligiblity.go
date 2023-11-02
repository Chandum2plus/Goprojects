package main

import (
	"fmt"
)

/*
10. Write a C program to determine eligibility for admission to a professional course based on the following criteria: Go to the editor
Eligibility Criteria : Marks in Maths >=65 and Marks in Phy >=55 and Marks in Chem>=50 and Total in all three subject >=190 or Total in
Maths and Physics >=140 ------------------------------------- Input the marks obtained in Physics :65 Input the marks obtained in Chemistry
:51 Input the marks obtained in Mathematics :72 Total marks of Maths, Physics and Chemistry : 188 Total marks of Maths and Physics : 137 The candidate is not eligible.
Expected Output :
The candidate is not eligible for admission.
*/
func admissionCriteria() {
	var math, phy, chem float64
	for {
		fmt.Print("Enter the Physics's Marks - ")
		fmt.Scan(&phy)

		fmt.Print("Enter the Math's Marks - ")
		fmt.Scan(&math)

		fmt.Print("Enter the Chemistry's Marks - ")
		fmt.Scan(&chem)

		if math >= 65 && phy >= 55 && chem >= 51 {
			fmt.Println("You are Eligible for Admission ")

		} else if (math + phy + chem) >= 190 {
			fmt.Println("you are Eligible for Admission ")

		} else if (math + phy) >= 140 {
			fmt.Println("You are Eligible for Admission ")

		} else {
			fmt.Println("You are not Eligible for Admission")
		}
	}
}

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
func main() {
	admissionCriteria()
}
