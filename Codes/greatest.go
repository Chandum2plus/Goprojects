package main

import "fmt"

// find the greatest number using the logical & operator and the
// nested if else condition
func greatest() {
	var num1, num2, num3 int
	fmt.Print("Enter the first number - ")
	fmt.Scan(&num1)
	fmt.Print("Enter the second number - ")
	fmt.Scan(&num2)
	fmt.Print("Enter the third number - ")
	fmt.Scan(&num3)

	//if num1 > num2 && num1 > num3 {
	//	fmt.Println(num1, " is greatest number ")
	//} else if num2 > num3 {
	//	fmt.Println(num2, " is Greatest number")
	//} else {
	//	fmt.Println(num3, " is Greatest number")
	//}

	//if num1 > num2 {
	//	fmt.Println(num1, " is greatest number ")
	//	if num2 > num1 {
	//		fmt.Println(num2, "Greatest number")
	//	}
	//}
	//if num1 > num3 {
	//	fmt.Println(num1, " Greatest Number")
	//	if num3 > num1 {
	//		fmt.Println(num3, "Greatest number")
	//	}
	//}
	//if num2 > num3 {
	//	fmt.Println(num2, " Greatest number")
	//	if num3 > num2 {
	//		fmt.Println(num3, " Greatest number")
	//	}
	//}

	if num1 > num2 {
		if num1 > num3 {
			fmt.Println(num1, " Greatest")
		} else {
			fmt.Println(num3, "Greatest")
		}
	} else if num2 > num3 {
		fmt.Println(num2, "Greatest")
	} else {
		fmt.Println(num3, "Greatest")
	}
	//
	//var state string
	//var bihar = "Bihar"
	//var up = "Up"
	//
	//fmt.Print("Enter the state - ")
	//fmt.Scan(&state)
	//
	//if state == bihar || state == up {
	//	fmt.Println("you are eligible for BPSC")
	//} else {
	//	fmt.Println("you are not eligible for BPSC")
	//}
	//
	//if state == bihar {
	//	fmt.Println("Your are eligible for BPSC")
	//} else if state == up {
	//	fmt.Println("you are eligible for BPSC ")
	//} else {
	//	fmt.Println("You are not eligible for BPSC ")
}

func main() {
	greatest()
}
