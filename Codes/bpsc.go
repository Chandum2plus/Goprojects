package main

import (
	"fmt"
)

// find the eligibility of state who can sit in the exam hall of the BPSC
func eligibility() {
	var state string

	for {
		fmt.Print("Enter the state - ")
		fmt.Scan(&state)

		// this is the first approach testing the Eligibility using the OR operator

		if state == "Bihar" || state == "Up" || state == "Jharkhand" {
			fmt.Println("You are Eligible for BPSC Exam")
		} else {
			fmt.Println("You are Not Eligible for BPSC Exam")
		}

		// This is the second approach testing the eligibility using the ladder if else

		if state == "Bihar" {
			fmt.Println("You are eligible for BPSC")
		} else if state == "Up" {
			fmt.Println("You are eligible for BPSC")
		} else if state == "jharkhand" {
			fmt.Println("You are eligible for BPSC")
		} else {
			fmt.Println("You are not eligible for BPSC")
		}
	}
}

// find the Greatest number among the four number
func greatests() {

	var num1, num2, num3, num4 int
	for {
		fmt.Print("Enter the first number - ")
		fmt.Scan(&num1)
		fmt.Print("Enter the second number - ")
		fmt.Scan(&num2)
		fmt.Print("Enter the third number - ")
		fmt.Scan(&num3)
		fmt.Print("Enter the fourth number - ")
		fmt.Scan(&num4)

		// This is the first approach using the logical & operator

		/*
			if num1 > num2 && num1 > num3 && num1 > num4 {
				fmt.Println(num1, " Greatest")
			} else if num2 > num3 && num2 > num4 {
				fmt.Println(num2, " Greatest")
			} else if num3 > num4 {
				fmt.Println(num3, "Greatest")
			} else {
				fmt.Println(num4, "Greatest")
			}
		*/

		// This is the second approach using the nested if else and without use any logical operator

		//if num1 > num2 {
		//	if num1 > num3 {
		//		if num1 > num4 {
		//			fmt.Println(num1, " Greatest")
		//		} else {
		//			fmt.Println(num4, " Greatest")
		//		}
		//
		//	}
		//} else if num3 > num2 {
		//	if num3 > num4 {
		//		fmt.Println(num3, " Greatest")
		//	} else {
		//		fmt.Println(num4, " Greatest")
		//	}
		//} else if num2 > num3 {
		//	if num2 > num4 {
		//		fmt.Println(num2, " Greatest")
		//	} else {
		//		fmt.Println(num4, " Greatest")
		//	}
		//}

		// This is the third approach
		if num1 > num2 {
			if num1 > num3 {
				if num1 > num4 {
					fmt.Println(num1, " Greatest")
				} else {
					fmt.Println(num4, " Greatest")
				}
			}
		} else if num2 > num3 {
			if num2 > num4 {
				fmt.Println(num2, " Greatest")
			} else {
				fmt.Println(num4, " Greatest")
			}

		} else if num3 > num2 {
			if num3 > num4 {
				fmt.Println(num3, " Greatest")
			} else {
				fmt.Println(num4, " Greatest")
			}
		}
	}
}

// Reverse of the Array
func rev() {
	// this  is the first approach using the user input
	var size int
	fmt.Print("Enter the size of the Array - ")
	fmt.Scan(&size)
	var arr = make([]int, size)
	fmt.Print("Enter the element of the Array - ")
	for i := 0; i < size; i++ {
		fmt.Scan(&arr[i])
		fmt.Print(" ", arr[i])
	}
	fmt.Print("\nReverse")
	for i := size - 1; i >= 0; i-- {
		fmt.Print(" ", arr[i])
	}

	// this is the second approach using the given value
	ar := [5]int{10, 20, 30, 40, 50}
	i := 0
	fmt.Print("Original Array =")
	for i = 0; i < 5; i++ {
		fmt.Print(" ", ar[i])
	}
	fmt.Print("\nReverse = ")
	for i = 4; i >= 0; i-- {
		fmt.Print(" ", ar[i])
	}
}

func main() {
	//eligibility()
	greatests()
	//rev()
}
