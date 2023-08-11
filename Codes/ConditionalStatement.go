// Program on Decision making

package main

import (
	"fmt"
)

// accept the two integer number and check whether they are eqaual or not
func Test() {
	var num float64
	var num2 float64
	fmt.Println("=== Here, WE WILL CHECK THE NUMBER WHETHER THEY ARE EQUAL OR NOT ===")
	fmt.Println()
	fmt.Print("\nEnter the First number -")
	fmt.Scan(&num)
	fmt.Print("\nEnter the Second number -")
	fmt.Scan(&num2)
	if num == num2 {
		fmt.Printf("%f = %f Equal", num, num2)
	} else {
		fmt.Printf("%f = %f Not Equal", num, num2)
	}

}

/*
	Write a  program to check whether a given number is even or odd.

Test Data : 15
Expected Output :
15 is an odd integer
*/
func evenOrOdd() {
	var num int
	fmt.Println("== CHECK THE WHETHER A NUMBER IS EVEN OR ODD  ==")
	fmt.Print("\nEnter the number -")
	fmt.Scan(&num)
	if num%2 == 0 {
		fmt.Printf("%d is Even number ", num)
	} else {
		fmt.Printf("%d is Odd number ", num)
	}

}

/*
Write a  program to check whether a given number is positive or negative.
Test Data : 15
Expected Output :
15 is a positive number
*/
func positive() {
	var num float64
	fmt.Println("== CHECK THE WHETHER NUMBER POSITIVE OR NEGATIVE ==")
	fmt.Print("\nEnter the number -")
	fmt.Scan(&num)
	if num > 0 {
		fmt.Printf("%f is Positive", num)
	} else {
		fmt.Printf("%f is Negative", num)
	}

}

/*
Write a  program to find whether a given year is a leap year or not.
Test Data : 2016
Expected Output :
2016 is a leap year.
*/
func yearLeapOrNot() {
	var year int
	fmt.Println("== FIND THE WHETHER YEAR IS LEAP YEAR OR NOT ==")
	fmt.Print("\nEnter the year - ")
	fmt.Scan(&year)
	if year%4 == 0 {
		fmt.Printf("%d-Leap Year", year)
	} else {
		fmt.Printf("%d-Not Leap Year", year)
	}

}

/*
Write a  program to read the age of a candidate and determine whether he is eligible to cast his/her own vote.
Test Data : 21
Expected Output :
Congratulation! You are eligible for casting your vote.
*/
func vote() {
	var age float64
	fmt.Println("== HERE,WE WILL CHECK YOUR ELIGIBLE FOR VOTE OR NOT ==")
	fmt.Print("\nEnter the Age - ")
	fmt.Scan(&age)
	if age >= 18 && age <= 80 {
		fmt.Println("You are eligible for Vote")

	} else if age < 18 {
		fmt.Println("You are not eligible for Vote")
	} else if age >= 80 {
		fmt.Println("You are Retired You can't Vote,so please go and rest ")
	}

}

/*
Write a C program to accept the height of a person in centimeters and categorize the person according to their height.
Test Data : 135
Expected Output :
The person is Dwarf.
*/
func menheight() {
	var height float64
	fmt.Println("== HERE, WE WILL CHECK THE MEN'S HEIGHT AVERAGE OR TALL ==")
	fmt.Print("\nEnter the men height in centimeter - ")
	fmt.Scan(&height)
	if height < 150 {
		fmt.Println("Dwarf Height")
	} else if height >= 150 && height < 165 {
		fmt.Println("Average Height")
	} else if height >= 165 && height <= 195 {
		fmt.Println("Tall Height")
	} else {
		fmt.Println("Abnormal Height")
	}

}

/*
Write a  program to find the largest of three numbers.
Test Data : 12 25 52
Expected Output :
1st Number = 12,        2nd Number = 25,        3rd Number = 52
The 3rd Number is the greatest among three
*/
func greatestNumb() {
	var num1, num2, num3 int
	fmt.Println("== FIND THE LARGEST OF THREE NUMBER ==")
	fmt.Print("\nEnter the first number - ")
	fmt.Scan(&num1)
	fmt.Print("\nEnter the second number - ")
	fmt.Scan(&num2)
	fmt.Print("\nEnter the third number - ")
	fmt.Scan(&num3)
	if num1 > num2 && num1 > num3 {
		fmt.Printf("%d is Greatest number \n", num1)
	} else if num2 > num1 && num2 > num3 {
		fmt.Printf("%d is Greatest number \n", num2)
	} else {
		fmt.Printf("%d is Greatest number \n", num3)
	}

}

/*
Write a C program to determine eligibility for admission to a professional course based on the following criteria: Go to the editor
Eligibility Criteria : Marks in Maths >=65 and Marks in Phy >=55 and Marks in Chem>=50 and Total in all three subject >=190 or Total in Maths and Physics >=140 -------------------------------------
Input the marks obtained in Physics :65
Input the marks obtained in Chemistry :51
Input the marks obtained in Mathematics :72
Total marks of Maths, Physics and Chemistry : 188
Total marks of Maths and Physics : 137 The candidate is not eligible.
Expected Output :
The candidate is not eligible for admission.
*/
func addmission() {
	var maths, phy, chem float64
	fmt.Println("== CHECK, YOUR ARE ELIGIBLE FOR COLLEGE ADMISSION OR NOT ==")
	fmt.Print("\nEnter the Math's marks - ")
	fmt.Scan(&maths)
	fmt.Print("\nEnter the Physics's marks - ")
	fmt.Scan(&phy)
	fmt.Print("\nEnter the Chemistry's marks - ")
	fmt.Scan(&chem)
	total := maths + phy + chem
	fmt.Println("Total Marks = ", total)
	if maths >= 65 && phy >= 55 && chem >= 50 {
		fmt.Println("You are Eligible for Admission")
	} else if (maths+phy+chem) >= 190 || (maths+phy) >= 140 {
		fmt.Println("You are Eligible for Admission")
	} else {
		fmt.Println("You are Not Eligible for Admission")
	}

}

/*
	Write a C program to read the roll no, name and marks of three subjects and calculate the total, percentage and division.

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
func percentatage() {
	var math, eng, sci, sst, sank float64
	var roll int
	var name string
	fmt.Println("== CHECK,YOUR PERCENTAGE AND DIVISION ==")
	fmt.Print("\nEnter Your Name - ")
	fmt.Scan(&name)
	fmt.Print("\nEnter Your Roll no. - ")
	fmt.Scan(&roll)
	fmt.Print("\nEnter the Math's Marks - ")
	fmt.Scan(&math)
	fmt.Print("\nEnter the English's Marks - ")
	fmt.Scan(&eng)
	fmt.Print("\nEnter the Science's Marks - ")
	fmt.Scan(&sci)
	fmt.Print("\nEnter the Social Science's Marks - ")
	fmt.Scan(&sst)
	fmt.Print("\nEnter the Sanskrit's Marks - ")
	fmt.Scan(&sank)
	fmt.Println()
	fmt.Println("======== RESULT SHOWN BELOW ========")
	fmt.Println("Name - ", name)
	fmt.Println("Roll - ", roll)
	if (math+eng+sst+sci+sank) >= 150 && (math+eng+sst+sci+sank) <= 224 {
		fmt.Println("Division - Third")
		var res float64
		res += (math + eng + sst + sci + sank)
		fmt.Println("Total Marks - ", res)
		perc := 0.0
		perc = (math + eng + sst + sci + sank) / 500 * 100
		fmt.Println("Percentage - ", perc)
	} else if (math+eng+sst+sci+sank) >= 225 && (math+eng+sst+sci+sank) <= 299 {
		fmt.Println("Division - Second ")
		var res float64
		res += (math + eng + sst + sci + sank)
		fmt.Println("Total Marks - ", res)
		perc := 0.0
		perc = (math + eng + sst + sci + sank) / 500 * 100
		fmt.Println("Percentage - ", perc)
	} else if (math+eng+sst+sci+sank) >= 300 && (math+eng+sst+sci+sank) <= 399 {
		fmt.Println("Division - First")
		var res float64
		res += (math + eng + sst + sci + sank)
		fmt.Println("Total Marks - ", res)
		perc := 0.0
		perc = (math + eng + sst + sci + sank) / 500 * 100
		fmt.Println("Percentage - ", perc)
	} else if (math+eng+sst+sci+sank) >= 400 && (math+eng+sst+sci+sank) <= 450 {
		fmt.Println("Division - District Top")
		var res float64
		res += (math + eng + sst + sci + sank)
		fmt.Println("Total Marks - ", res)
		perc := 0.0
		perc = (math + eng + sst + sci + sank) / 500 * 100
		fmt.Println("Percentage - ", perc)
	} else if (math+eng+sst+sci+sank) > 400 && (math+eng+sst+sci+sank) <= 500 {
		fmt.Println("Division - State Top")
		var res float64
		res += (math + eng + sst + sci + sank)
		fmt.Println("Total Marks - ", res)
		perc := 0.0
		perc = (math + eng + sst + sci + sank) / 500 * 100
		fmt.Println("Percentage - ", perc)
	} else if (math+eng+sst+sci+sank) >= 100 && (math+eng+sst+sci+sank) <= 149 {
		fmt.Println("Division - Failed")
		var res float64
		res += (math + eng + sst + sci + sank)
		fmt.Println("Total Marks - ", res)
		perc := 0.0
		perc = (math + eng + sst + sci + sank) / 500 * 100
		fmt.Println("Percentage - ", perc)
	}
}

// find the age from date of birth
func age() {
	var crDate, crMonth, crYear, brDate, brMonth, brYear int
	fmt.Println("Enter the Date of Birth like - / Month / Year")
	fmt.Scanln(&brDate, brMonth, brYear)
	// Day of the every month
	month := [12]int{31, 28, 31, 30, 31, 30,
		31, 31, 30, 31, 30, 31}

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

func main() {

	for {
		fmt.Println("\n")
		fmt.Println("=============== MENU BAR ==============")
		fmt.Println("||    1 - Number Comparison          ||")
		fmt.Println("||    2 - Test Even or Odd           ||")
		fmt.Println("||    3 - Test Positive/Negative     ||")
		fmt.Println("||    4 - Test Year Leap Or Not      ||")
		fmt.Println("||    5 - Test Eligible for Vot      ||")
		fmt.Println("||    6 - Test Men's Height          ||")
		fmt.Println("||    7 - Test Greatest Number       ||")
		fmt.Println("||    8 - Test Eligible for Adm.     ||")
		fmt.Println("||    9 - Test the Result of Student ||")
		fmt.Println("||   10 - Test the Age from DOB      ||")
		fmt.Println("=======================================")
		var choice int
		fmt.Print("\nEnter your Choice - ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			Test()
			break
		case 2:
			evenOrOdd()
			break
		case 3:
			positive()
			break
		case 4:
			yearLeapOrNot()
			break
		case 5:
			vote()
			break
		case 6:
			menheight()
			break
		case 7:
			greatestNumb()
			break
		case 8:
			addmission()
		case 9:
			percentatage()
		case 10:
			age()
		default:
			fmt.Println("Please Enter the valid Choice We don't have 11th Choice ")
		}
	}

}
