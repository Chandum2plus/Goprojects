package main

import "C"
import (
	"fmt"
	//"fmt"
)

func revers() {

	a := 'A'
	b := 'C'
	c := 'D'
	fmt.Printf("ascending = %c %c %c\n", a, b, c)
	fmt.Printf("reverse = %c %c %c\n", c, b, a)

}
func perimeter() {
	var length, breath float64
	fmt.Print("Enter the length - ")
	fmt.Scan(&length)
	fmt.Print("Enter the Breath - ")
	fmt.Scan(&breath)

	peri := 2 * (length + breath)
	fmt.Println("Perimeter of rectangle -", peri)
	area := length * breath
	fmt.Println("Area - ", area)
}
func years() {
	for {
		var day, year, week int

		fmt.Print("\nEnter the Days - ")
		fmt.Scan(&day)
		year = day / 365

		week = (day % 365) / 7
		day = day - ((year * 365) + (week * 7))
		fmt.Printf("year = %d\n", year)

		fmt.Printf("week = %d\n", week)
		fmt.Printf("day = %d", day)
	}

}

// find the product of a number
func products() {
	a := 25
	b := 15
	fmt.Print("first number -", a)
	fmt.Print("\nsecond number - ", b)
	product := a * b
	fmt.Print("\nproduct of the number =", product)
}

// find the average of weight of item
func avg() {
	var weight1 float64
	var item1 float64
	var weight2 float64
	var item2 float64
	fmt.Print("Enter the weight of first item -")
	fmt.Scan(&weight1)
	fmt.Print("Enter number of the first item -")
	fmt.Scan(&item1)
	fmt.Print("Enter the weight of second item -")
	fmt.Scan(&weight2)
	fmt.Print("Enter the number of second item -")
	fmt.Scan(&item2)
	firstTotalItem := weight1 * item1
	fmt.Print("\nFirst total item weight =", firstTotalItem)
	secondTotalItem := weight2 * item2
	fmt.Print("\nSecond total item weight =", secondTotalItem)
	avrege := ((weight1 * item1) + (weight2 * item2)) / (item2 + item1)
	result := (secondTotalItem + firstTotalItem) / (item2 + item1)
	fmt.Print("\nAverage =", avrege)
	fmt.Print("\nResult =", result)
}

/*
	Write a C program that accepts an employee's ID, total worked hours in a month and the amount he received per hour. Print the ID and salary (with two decimal places) of the employee for a particular month.

Test Data :
Input the Employees ID(Max. 10 chars): 0342
Input the working hrs: 8
Salary amount/hr: 15000
Expected Output:
Employees ID = 0342
Salary = U$ 120000.00
*/
func employee() {
	var empId int
	var hours float64
	var salary float64
	fmt.Print("Enter the employee ID - ")
	fmt.Scan(&empId)
	fmt.Print("Enter the salary amount/hr - ")
	fmt.Scan(&salary)
	fmt.Print("Enter the working hours - ")
	fmt.Scan(&hours)
	Result := salary * hours
	fmt.Print("\nEmployee Id -", empId)
	fmt.Print("\nEmployee salary is = ", Result)

}

/*
13. Write a C program that accepts three integers and finds the maximum of three.
Test Data :
Input the first integer: 25
Input the second integer: 35
Input the third integer: 15
Expected Output:
Maximum value of three integers: 35
*/
func maxNumber() {
	for {
		var num1, num2, num3 int
		fmt.Print("\nEnter the first number - ")
		fmt.Scan(&num1)
		fmt.Print("\nEnter the second number - ")
		fmt.Scan(&num2)
		fmt.Print("\nEnter the third number - ")
		fmt.Scan(&num3)
		if num1 > num2 && num1 > num3 {
			fmt.Print("\nfirst number is greatest number - ", num1)
		} else if num2 > num1 && num2 > num3 {
			fmt.Print("\n Second number is greatest number - ", num2)
		} else {
			fmt.Print("\nThird number is greatest number - ", num3)
		}
	}
}

/*
14. Write a C program to calculate a bike’s average consumption from the given total distance (integer value) travelled (in km) and spent fuel (in litters, float number – 2 decimal points).
Test Data :
Input total distance in km: 350
Input total fuel spent in liters: 5
Expected Output:
Average consumption (km/lt) 70.000
*/
func bike() {
	for {
		var distance float64
		var fuel float64
		fmt.Print("\nEnter the Total distance in km -")
		fmt.Scan(&distance)
		fmt.Print("\nEnter the Total fuel spent in liters - ")
		fmt.Scan(&fuel)
		result := distance / fuel
		fmt.Print("\nAverage consumption (km/ltr) = ", result)
	}
}

/*
Write a C program to convert a given integer (in seconds) to hours, minutes and seconds.
Test Data :
Input seconds: 25300
Expected Output:
There are:
H:M:S - 7:1:40
*/
func times() {
	var sec, h, m, s int
	fmt.Print("Enter the second - ")
	fmt.Scan(&sec)
	h = (sec / 3600)
	m = (sec - (3600 * h)) / 60
	s = (sec - (3600 * h) - (m * 60))
	fmt.Printf("\nH:M:S - %d :%d :%d", h, m, s)

}

// How to concatenation of a string using plus (+) operator
func concatenating() {
	//str := "Chandu"
	//str1 := "Kumar"
	//fmt.Print("result =", str, " "+str1)
	/*
		var a int
		var b string
		var c float64

	*/
	var a, b, c = 20, "Chandu", 40
	fmt.Printf("value of a - %d\n", a)
	fmt.Printf("value of b - %s\n", b)
	fmt.Printf("value of c - %d\n", c)
	//d, r, e := 2, 3.4, "Chand"
	//fmt.Printf("value of d = %d value of r = %f and value of e = %s\n", d, r, e)
}
func main() {
	//revers()
	//perimeter()
	//years()
	//products()
	//avg()
	//employee()
	//maxNumber()
	//bike()
	//	times()
	concatenating()
}
