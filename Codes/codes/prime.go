package main

import "fmt"

// find the prime number
func prime() {
	var num int
	fmt.Print("enter the number - ")
	fmt.Scan(&num)
	count := 0

	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			count++
			break
		}
	}
	if count == 0 && count != 1 {
		fmt.Println(num, "Prime numbe ")
	} else {
		fmt.Println(num, "Not prime number")
	}
}

// find the prime between the min and max
func primeMinMax() {
	var num, min, max int
	count := 0
	fmt.Print("enter the min and max number - ")
	fmt.Scanln(&min, &max)
	fmt.Println("Prime Number between ", min, "and ", max, "are -")
	for num = min; num <= max; num++ {
		count++
		for i := 2; i < num/2; i++ {
			if num%i == 0 {
				count++
				break
			}
		}
		if count == 0 && num != 1 {
			fmt.Print(num, "\t")
		}
	}
	fmt.Println()

	// var primNum, primMin, primMax, primcount int

	// fmt.Print("Enter the Minimum and Maximum Limit for Prime Numbers = ")
	// fmt.Scanln(&primMin, &primMax)

	// fmt.Println("Prime Numbers between ", primMin, " and ", primMax, " are ")
	// for primNum = primMin; primNum <= primMax; primNum++ {
	// 	primcount = 0
	// 	for i := 2; i < primNum/2; i++ {
	// 		if primNum%i == 0 {
	// 			primcount++
	// 			break
	// 		}
	// 	}
	// 	if primcount == 0 && primNum != 1 {
	// 		fmt.Print(primNum, "\t")
	// 	}
	// }
	// fmt.Println()

}
func main() {

	//prime()
	primeMinMax()

}
