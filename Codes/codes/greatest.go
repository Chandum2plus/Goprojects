package main

import "fmt"

// Find the Greatest number among the five number
func greatest() {
	var num1, num2, num3, num4, num5, gret int
	fmt.Print("Enter the first number - ")
	fmt.Scan(&num1)
	fmt.Print("Enter the second number - ")
	fmt.Scan(&num2)
	fmt.Print("Enter the third number - ")
	fmt.Scan(&num3)
	fmt.Print("Enter the fourth number - ")
	fmt.Scan(&num4)
	fmt.Print("Enter the fifth number - ")
	fmt.Scan(&num5)

	// checking the condition
	// if num1 will be greater than the num2 then enter the nested if part otherwise it will
	// terminate the condition
	if num1 > num2 {
		if num1 > num3 {
			if num1 > num4 {
				if num1 > num5 {
					gret = num1 //fmt.Println(num1, "= Greatest") //  if the num1 greater than all number then print num1 Greatest
				} else { // otherwise print the else part
					//fmt.Println(num5, "= Greatest")
					gret = num5
				}
			}
		}
	} else if num2 > num1 {
		if num2 > num3 {
			if num2 > num4 {
				if num2 > num5 {
					gret = num2
					//fmt.Println(num2, "= Greatest")
				} else {
					gret = num5
					//fmt.Println(num5, "= Greatest")
				}
			}
		}
	} else if num3 > num1 {
		if num3 > num2 {
			if num3 > num4 {
				if num3 > num5 {
					//	fmt.Println(num3, "= Greatest")
					gret = num3

				} else {
					//fmt.Println(num5, "= Greatest")
					gret = num5
				}
			}
		}
	} else if num4 > num1 {
		if num4 > num2 {
			if num4 > num3 {
				if num4 > num5 {
					gret = num4
					//fmt.Println(num4, "= Greatest")
				} else {
					gret = num5
					//fmt.Println(num5, "= Greatest")
				}
			}
		}
	}
	fmt.Println("Greatest Number - ", gret)
}
func main() {
	greatest()
}
