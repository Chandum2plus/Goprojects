package main

import (
	"fmt"
)

// Testing the number positive or negative
func pos() {
	var num int
	fmt.Print("enter the number - ")
	fmt.Scan(&num)
	if num > 0 {
		fmt.Println(num, "- Positive ")
	} else if num < 0 {
		fmt.Println(num, "- Negative")
	} else {
		fmt.Println(num, "- Zero")
	}
}

// Testing the Age young,adult or kid
func ages() {
	var age float64
	fmt.Print("enter your age - ")
	fmt.Scan(&age)
	if age < 20 {
		fmt.Println("You are kid")
	} else if age > 50 {
		fmt.Println("You are old")
	} else if age >= 20 && age <= 35 {
		fmt.Println("ekdam sudh jawan")
	}
	a := 9
	if a == 9 {
		fmt.Println("Nine")
	}
	if a > 0 {
		fmt.Println("Positive")
	} else {
		fmt.Println("Negative")
	}

}
func testing() {
	//fmt.Println("The List of alphabet")
	//for ch := 'A'; ch <= 'Z'; ch++ {
	//	fmt.Printf("%c ", ch)
	//}
	//fmt.Println()

	var val rune
	fmt.Print("Enter the value - ")
	fmt.Scan(&val)

	if val == 'a' || val == 'z' {
		fmt.Printf("%c - This is lower case", val)
	} else if val == 'A' || val == 'Z' {
		fmt.Printf("%c - This is Upper case", val)
	}

}

func even() {
	var num int
	fmt.Print("enter the number -")
	fmt.Scan(&num)
	i := 1
	sum := 0
	for i <= num {
		if i%2 == 0 {
			fmt.Println(i)
			sum += i
		}
		i++
	}
	fmt.Println("sum =", sum)
}
func pattern() {
	var num int
	fmt.Print("Enter the number - ")
	fmt.Scan(&num)
	i := 1
	for i <= num {
		j := 1
		for j <= num {
			fmt.Print(" ", j)
			j++
		}
		fmt.Println()
		i++
	}
}
func pattern1() {
	var row int
	fmt.Print("enter the row - ")
	fmt.Scan(&row)
	i := 1
	j := 1
	k := 1
	for i = 1; i <= row; i++ {
		for j = 1; j <= i; j++ {
			fmt.Print(" ", k)
			k++
			//fmt.Println()
		}
		i++
		fmt.Println()
	}
}
func main() {
	//pos()
	//ages()
	//testing()
	even()
	//pattern()
	//pattern1()

}
