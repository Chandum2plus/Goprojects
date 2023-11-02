package main

import "fmt"

/*
4. Write a C program to find whether a given year is a leap year or not.
Test Data : 2016
Expected Output :
2016 is a leap year.
*/

// check the Year is leap or not
func leap() {
	var year int
	for {
		fmt.Print("\nEnter the Year - ")
		fmt.Scan(&year)
		if year%400 == 0 {
			fmt.Printf("%d is Leap Year", year)
		} else if year%100 == 0 {
			fmt.Printf("%d is Leap Year", year)

		} else if year%4 == 0 {
			fmt.Printf("%d is Leap Year", year)
		} else {
			fmt.Printf("%d is Not Leap Year", year)

		}
	}
}
func main() {
	leap()
}
