package main

import (
	"fmt"
)

// Display the first 10 natural number
func naturalNumber10() {
	for i := 1; i <= 10; i++ {
		fmt.Print(" ", i)
	}
}

// sum of the 10 natural number
func sumNatural() {
	sum := 0
	for i := 1; i <= 10; i++ {
		fmt.Print(" ", i)
		sum += i
	}
	fmt.Println("\nsum = ", sum)
}

// Display n term of the number and their sum
func nTermSum() {
	var num, sum int
	fmt.Print("Enter the Number - ")
	fmt.Scan(&num)

	for i := 1; i <= num; i++ {
		fmt.Print(" ", i)
		sum += i
	}
	fmt.Printf("\nSum of number upto %d Terms = %d", num, sum)
}

/*
4. Write a program in C to read 10 numbers from the keyboard and find their sum and average.
Test Data :
Input the 10 numbers :
Number-1 :2
...
Number-10 :2
Expected Output :
The sum of 10 no is : 55
The Average is : 5.500000
*/
func arrSum() {

	// This is the first approach using the Array
	/*
		var size, sum, avg int
		fmt.Print("Enter the size of the Array - ")
		fmt.Scan(&size)
		fmt.Print("Enter the Element\n")
		arr := make([]int, size)
		for i := 0; i < size; i++ {
			fmt.Printf("Number %d -", i)
			fmt.Scan(&arr[i])
		}
		for i := 0; i < size; i++ {
			sum += arr[i]
		}
		avg = sum / size
		fmt.Printf("\nSum = %d", sum)
		fmt.Printf("\nAverage = %d", avg)

	*/

	// This  is the second Approach
	var s, n, av float64
	fmt.Print("Enter the 10 Number \n")
	for i := 1; i <= 10; i++ {
		fmt.Printf("Number %d - ", i)
		fmt.Scan(&n)
		s += n
	}
	av = s / n
	fmt.Printf("\nsum of %f terms = %f", n, s)
	fmt.Printf("\nAverage = %f", av)
}

/*
5. Write a program in C to display the cube of the number up to an integer.
Test Data :
Input number of terms : 5
Expected Output :
Number is : 1 and cube of the 1 is :1
Number is : 2 and cube of the 2 is :8
Number is : 3 and cube of the 3 is :27
Number is : 4 and cube of the 4 is :64
Number is : 5 and cube of the 5 is :125
*/
func cube() {
	var num int
	fmt.Print("Enter the Number - ")
	fmt.Scan(&num)
	for i := 1; i <= num; i++ {

		fmt.Printf("cube %d : and Cube of %d = %d\n", i, i, (i * i * i))
	}
}

/*
6. Write a program in C to display the multiplication table for a given integer.
Test Data :
Input the number (Table to be calculated) : 15
Expected Output :
15 X 1 = 15
...
...
15 X 10 = 150
*/

func table() {
	var num int
	fmt.Print("Enter the Number - ")
	fmt.Scan(&num)
	for i := 1; i <= 10; i++ {
		//	multi=num*i
		fmt.Printf("%d X %d = %d\n", num, i, num*i)
	}
}

/*
7. Write a program in C to display the multiplier table vertically from 1 to n.
Test Data :
Input upto the table number starting from 1 : 8
Expected Output :
Multiplication table from 1 to 8
1x1 = 1, 2x1 = 2, 3x1 = 3, 4x1 = 4, 5x1 = 5, 6x1 = 6, 7x1 = 7, 8x1 = 8
...
1x10 = 10, 2x10 = 20, 3x10 = 30, 4x10 = 40, 5x10 = 50, 6x10 = 60, 7x10 = 70, 8x10 = 80
*/
func tables() {
	var num int
	fmt.Print("Enter the Number - ")
	fmt.Scan(&num)
	for i := 1; i <= 10; i++ {
		for j := 1; j <= num; j++ {
			fmt.Printf("%d X %d = %d ,", j, i, i*j)
		}
		fmt.Println()
	}

}

/*
8. Write a C program to display the n terms of odd natural numbers and their sum.
Test Data
Input number of terms : 10
Expected Output :
The odd numbers are :1 3 5 7 9 11 13 15 17 19
The Sum of odd Natural Number upto 10 terms : 100
*/
func oddSums() {
	var num, sum int
	fmt.Print("Enter the Number - ")
	fmt.Scan(&num)
	count := 0
	fmt.Print("Odd Numbers = ")
	for i := 1; i <= num; i++ {
		fmt.Printf(" %d", 2*i-1)
		count++
		sum += 2*i - 1
		//if i%2 != 0 {
		//	fmt.Print(" ", i)
		//	sum += i
		//}
	}
	fmt.Printf("\nTotal Odd = %d", count)
	fmt.Printf("\nsum of the Odd %d terms =%d ", num, sum)
}

/*
9. Write a program in C to display a pattern like a right angle triangle using an asterisk.
The pattern like :
*
**
***
****
*/
func pattern() {
	for i := 1; i <= 4; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")

		}
		fmt.Print("\n")
	}
}

/*
10. Write a C program to display a pattern like a right angle triangle with a number.

The pattern like :

1
12
123
1234
*/
func paternNumber() {
	for i := 1; i <= 4; i++ {
		for j := 1; j < i; j++ {
			fmt.Print(j)
		}
		fmt.Println(i)
	}
}

/*
11. Write a program in C to make such a pattern like a right angle triangle with a number which will repeat a number in a row.
The pattern like :

	1
	22
	333
	4444
*/
func traingleNumber() {

	for i := 1; i <= 4; i++ {
		for j := 1; j <= i; j++ {

			fmt.Print(i)
		}
		fmt.Println()
	}
}

/*
12. Write a program in C to make such a pattern like a right angle triangle with the number increased by 1.
The pattern like :

	1
	2 3
	4 5 6
	7 8 9 10
*/
func numberIncrease() {
	k := 1
	for i := 1; i <= 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(" ", k)
			k++
		}
		fmt.Println()
	}
}
func main() {
	//naturalNumber10() // print the 10 natural number
	//sumNatural() // Print the sum of the 10 natural number
	//nTermSum() // Print the sum of N term of the number
	//arrSum()  // sum of the Array and their average
	// cube() // Print the cube of elements
	//table()   //Print the table
	tables() // Print the table of the input number
	//oddSums()   // Print the Odd sum of
	//pattern()   // this is print the pattern
	//	paternNumber() // this is print the  pattern of number
	//traingleNumber() // This is print the repeat number pattern
	//numberIncrease()

}
