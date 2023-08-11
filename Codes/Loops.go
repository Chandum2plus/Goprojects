package main

import (
	"fmt"
)

func forloop() {
	fmt.Println("Display the first 10 Natural Number -")
	for i := 1; i <= 10; i++ {
		fmt.Print(" ", i)
	}
}
func forInfinite() {
	var res = "Chandu"

	for {
		fmt.Println(res)
	}
}
func forAswhileLoop() {
	fmt.Println("First as a while loop -")
	i := 0
	for i < 3 {
		i += 2

	}

	fmt.Println("Second as a while loop -")
	fmt.Println("Result - ", i)
	number := 1
	fmt.Println("Second Result")
	for number <= 5 {
		fmt.Println(number)
		number++
	}

	fmt.Println("Third as a while loop -")
	multiplier := 1
	fmt.Println("Product of 5 ")
	for multiplier <= 10 {
		product := 5 * multiplier
		fmt.Printf("5 * %d = %d\n", multiplier, product)
		multiplier++
	}

}
func forAsDoWhileLoop() {
	number := 1 // initialized the number by 1
	// Loop,that runs infinitely
	for {
		// Condition to terminate the loop
		if number > 10 { // Here, condition will true till 10
			break // it is used to terminate the loop when it will greater than 10
		}
		fmt.Println(number)
		number++ // Here, it increament the number by 1
	}
}

func forRangeLoop() {

	// This is the first syntax of the array declaration
	var strar = [5]string{"Chandu", "kumar", "aditya", "kumar", "M2-Plus"}
	fmt.Println("Displaying the String Array Using Range loop -")
	for i, j := range strar {
		fmt.Println(i, j)
	}

	// This is the Second syntax of the array declaration
	fmt.Println("Displaying the Integer Array Using Range Loop ")
	var ar = [...]int{11, 22, 33, 44, 45, 66, 57}
	fmt.Println("Index  Values")
	for i, j := range ar {
		fmt.Printf(" %d\t%d\n", i, j)
	}
}
func main() {

	forRangeLoop()
	//forAsDoWhileLoop()
	//forAswhileLoop()
	//forInfinite()
	//forloop()
}
