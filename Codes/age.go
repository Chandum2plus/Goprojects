package main

import (
	"fmt"
	"strconv"
)

// find the age from date of birth
func ageInDay() {
	var crDate, brDate string
	fmt.Println("Enter the Date of Birth like - / Month / Year")
	fmt.Scanln(&brDate)
	fmt.Println("Enter the current date -")
	fmt.Scanln(&crDate)
	fmt.Println("cr Date =", crDate)
	fmt.Println("br moth =", brDate)
	//var Rescr,Resbr int
	_, Rescr := strconv.Atoi(crDate)
	_, Resbr := strconv.Atoi(brDate)
	fmt.Printf("Rescr = %d\n", Rescr, " Resbr  = %d", Resbr)

	// Day of the every month
	//month := [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	// if the date is greater than current date ,then do not count this
	// month and add 30 to the date so as to subtract the date and get the remaining days

	//if brDate > crDate {
	//	crDate = crDate + month[brMonth]
	//	crMonth = crMonth + 1
	//}
	//
	//// if the birth month exceeds current month,then do not count this
	//// year and add 12 to the month so that we can subtract
	//// and find out the difference
	//
	//if brMonth > crMonth {
	//	crYear = crYear - 1
	//	crMonth = crMonth - 12
	//}
	////if brYear>crYear{
	////	month
	////}
	//// calculated date month year
	//resDate := crDate - brDate
	//resMonth := crMonth - brMonth
	//resYear := crYear - brYear

	// display the result
	//fmt.Printf("\nAge = \nYear: %d  Month: %d  Days: %d\n", resYear, resMonth, resDate)

}
func main() {
	ageInDay()
}
