package main

import "fmt"

/*
5. Write a C program to read the age of a candidate and determine whether he is eligible to cast his/her own vote.
Test Data : 21
Expected Output :
Congratulation! You are eligible for casting your vote.
*/
func vote(year float64) {
	if year >= 18 {
		fmt.Println("Congratulation ! You are Eligible for vote ")
	} else {
		fmt.Println("Sorry ! You are not Eligible for vote ")
	}
}
func main() {
	var y float64
	fmt.Print("Enter your Age - ")
	fmt.Scan(&y)
	vote(y)

}
