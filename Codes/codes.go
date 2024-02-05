package main

import (
	"fmt"
)

/*
18. Write a C program to calculate profit and loss on a transaction.
Test Data :
500 700
Expected Output :
You can booked your profit amount : 200
*/
func profitAndLose() {
	var sellingPrice, costPrice, Profit_Lose float64
	fmt.Print("\nEnter the Cost Price - ")
	fmt.Scan(&costPrice)
	fmt.Print("Enter the Selling Price - ")
	fmt.Scan(&sellingPrice)
	if sellingPrice > costPrice {
		Profit_Lose = sellingPrice - costPrice
		fmt.Printf("Your Profit Amount - %f", Profit_Lose)
	} else if costPrice > sellingPrice {
		Profit_Lose = costPrice - sellingPrice
		fmt.Printf("Your Lose Amount - %f", Profit_Lose)
	} else {
		fmt.Printf("No Profit and No lose Running Condition - %f", Profit_Lose)
	}
}

/*
19. Write a program in C to calculate and print the electricity bill of a given customer. The customer ID, name, and unit consumed by
the user should be captured from the keyboard to display the total amount to be paid to the customer.

The charge are as follow :

	float chg, surchg=0, gramt,netamt;

Unit	Charge/unit
upto 199	@1.20
200 and above but less than 400	@1.50
400 and above but less than 600	@1.80
600 and above	@2.00
If bill exceeds Rs. 400 then a surcharge of 15% will be charged and the minimum bill should be of Rs. 100/-

Test Data :
1001
James
800
Expected Output :
Customer IDNO :1001
Customer Name :James
unit Consumed :800
Amount Charges @Rs. 2.00 per unit : 1600.00
Surchage Amount : 240.00
Net Amount Paid By the Customer : 1840.00
*/
func electricityBill() {
	var name string
	var consumerID int
	var consumed, charge, surcharge, gradeAmount, netAmount float64
	fmt.Print("Enter Name of the Costumer - ")
	fmt.Scan(&name)
	fmt.Print("Enter the Consumer ID - ")
	fmt.Scan(&consumerID)
	fmt.Print("Enter Consumed by the Costumer - ")
	fmt.Scan(&consumed)
	if consumed < 200 {
		charge = 1.20
	} else if consumed >= 200 && consumed < 400 {
		charge = 1.50
	} else if consumed >= 400 && consumed < 600 {
		charge = 1.80
	} else {
		charge = 2.0
	}
	gradeAmount = consumed * charge
	if gradeAmount > 300 {
		surcharge = gradeAmount * 15 / 100
		netAmount = gradeAmount + surcharge
	}
	if netAmount < 100 {
		netAmount = 100
	}
	fmt.Println("\n========= RESULT ==========")
	fmt.Println("Consumer Name - ", name)
	fmt.Println("Consumer ID - ", consumerID)
	fmt.Println("Consumed - ", consumed)
	fmt.Println("Charge/Unit - ", charge)
	fmt.Println("Surcharge - ", surcharge)
	fmt.Printf("Grade Amount @Rs.%4.2f per Unit - %8.2f\n ", charge, gradeAmount)
	fmt.Println("Net Amount - ", netAmount)
}

/*
21. Write a C program to read any day number in integer and display the day name in word format.
Test Data :
4
Expected Output :
Thursday
*/
func day() {
	var choice int
	fmt.Print("Enter the choice - ")
	fmt.Scan(&choice)
	res := ""
	switch choice {
	case 1:
		res = "Sunday"
	case 2:
		res = "Monday"
	case 3:
		res = "Tuesday"
	case 4:
		res = "Wednesday"
	case 5:
		res = "Thursday"
	case 6:
		res = "Friday"
	case 7:
		res = "Saturday"
	default:
		res = "Invalid Choice"
	}
	fmt.Println(res)
}

/*
24. Write a program in C to read any Month Number in integer and display the number of days for this month.
Test Data :
7
Expected Output :
Month have 31 days
*/
func monthDay() {
	var month string
	for {
		fmt.Print("Enter the Month - ")
		fmt.Scan(&month)
		switch month {
		case "January", "March", "May", "July", "August", "October", "December":
			fmt.Println("Month have 31 days")
			break
		case "February":
			fmt.Println("Month has 28 day")
			fmt.Println("In leap year month has 29 days")
			var year int
			fmt.Println("CHECK THE YEAR IS LEAP OR NOT")
			fmt.Print("Enter the Year - ")
			fmt.Scan(&year)
			if year%4 == 0 {
				fmt.Printf("%d - Leap Year\n", year)
			} else {
				fmt.Printf("%d - Not Leap Year\n", year)
			}
			break

		case "April", "June", "September", "November":
			fmt.Println("Month has 30 days")
			break
		default:
			fmt.Println("Invalid Month Please Enter the Valid Month")

		}
	}
}

/*
26. Write a program in C which is a Menu-Driven Program to perform a simple calculation.
Test Data :
10
2
3
Expected Output :
The Multiplication of 10 and 2 is: 20
*/
func menuDriven() {
	var num1, num2 float64
	for {
		fmt.Print("\nEnter the First Number - ")
		fmt.Scan(&num1)
		fmt.Print("Enter the Second Number - ")
		fmt.Scan(&num2)
		var opt int

		fmt.Print("1- Addition\n2- Subtraction\n3- Multiplication\n4- Division\n")
		fmt.Print("Enter the Option - ")
		fmt.Scan(&opt)
		switch opt {
		case 1:
			fmt.Printf("Addition %f + %f = %f\n ", num1, num2, num1+num2)
			break
		case 2:
			fmt.Printf("Subtraction %f - %f = %f\n", num1, num2, num1-num2)
			break
		case 3:
			fmt.Printf("Multiplication %f x %f = %f\n", num1, num2, num1*num2)
			break
		case 4:
			if num2 == 0 {
				fmt.Println("Second Value is Divide by Zero")
				fmt.Printf("Divide by Zero %f / %f = %f ", num1, num2, num1/num2)
			} else {
				fmt.Printf("Division %f / %f = %f\n ", num1, num2, num1/num2)
			}
			break
		default:
			fmt.Println("Please Enter the valid choice")

		}
	}
}
func main() {
	//menuDriven()
	//monthDay()
	//	day()
	//electricityBill()
	//profitAndLose()
}
