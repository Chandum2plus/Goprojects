package main

import (
	"fmt"
)

// example of the basic arithmetic operation in golang  .........

// sum of the two numbers ...
func sum(num1, num2 int) int {
	res := num1 + num2
	return res
}

// subtraction of the two number ........
func subtraction(num1, num2 int) int {
	sub := num1 - num2
	return sub

}

// multiplication of two number
func multiplication(num1, num2 int) int {
	multi := num1 * num2
	return multi
}

// division of two number
func division(num1, num2 int) int {
	div := num1 / num2
	return div
}

// modulus of two number
func modulus(num1, num2 int) int {
	mod := num1 % num2
	return mod
}
func main() {
	var num1, num2 int
	var choice int
	for {
		fmt.Println()
		fmt.Println("=========== MENU BAR ============")
		fmt.Println("=      1- sum of two number        =")
		fmt.Println("=      2- sub of two number        =")
		fmt.Println("=      3- multi of two number      =")
		fmt.Println("=      4- div of two number        =")
		fmt.Println("=      5- Mod of two number        =")
		fmt.Println("====================================")
		fmt.Println()
		fmt.Println("Enter your choice  -")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Println("FOR SUM ")
			fmt.Println("Enter the first number -")
			fmt.Scanln(&num1)
			fmt.Println("Enter the second number -")
			fmt.Scanln(&num2)
			fmt.Println("sum =", sum(num1, num2))
			break
		case 2:
			fmt.Println("FOR SUBTRACTION")
			fmt.Println("Enter the first number -")
			fmt.Scanln(&num1)
			fmt.Println("Enter the second number -")
			fmt.Scanln(&num2)
			fmt.Println("subtraction =", subtraction(num1, num2))
			break
		case 3:
			fmt.Println("FOR MULTIPLICATION")
			fmt.Println("Enter the first number -")
			fmt.Scanln(&num1)
			fmt.Println("Enter the second number -")
			fmt.Scanln(&num2)
			fmt.Println("Multiplication =", multiplication(num1, num2))
			break
		case 4:
			fmt.Println("FOR DIVISION")
			fmt.Println("Enter the first number -")
			fmt.Scanln(&num1)
			fmt.Println("Enter the second number -")
			fmt.Scanln(&num2)
			fmt.Println("Division =", division(num1, num2))
			break
		case 5:
			fmt.Println("FOR MODULUS")
			fmt.Println("Enter the first number -")
			fmt.Scanln(&num1)
			fmt.Println("Enter the second number -")
			fmt.Scanln(&num2)
			fmt.Println("Modulus =", modulus(num1, num2))
			break
		default:
			fmt.Println("Please Enter the valid choice ")
		}

	}
}
