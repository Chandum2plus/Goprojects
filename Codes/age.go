package main

import "fmt"

func main() {
	ages()
}
func ages() {
	var crDate, crMonth, crYear, brDate, brMonth, brYear int
	fmt.Println("Enter the Date of Birth like - Date / Month / Year")
	fmt.Scanln(&brDate, &brMonth, &brYear)
	// Day of the every month
	month := [...]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	// if the date is greater than current date ,then do not count this
	// month and add 30 to the date so as to subtract the date and get the remaining days

	if brDate > crDate {
		crDate = crDate + month[brMonth-1]
		crMonth = crMonth - 1
	}

	// if the birth month exceeds current month,then do not count this
	// year and add 12 to the month so that we can subtract
	// and find out the difference

	if brMonth > crMonth {
		crYear = crYear - 1
		crMonth = crMonth + 12
	}
	// calculated date month year
	resDate := crDate - brDate
	resMonth := crMonth - brMonth
	resYear := crYear - brYear

	// display the result
	fmt.Printf("\nAge = \nYear:- %d  Month:- %d  Days:- %d\n", resYear, resMonth, resDate)
}
