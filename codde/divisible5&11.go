package main

import "fmt"

// find that number which is divisible by 5 and 11 ..............

func divisible() {
	var num int
	for {
		fmt.Print("\nenter the number - ")
		fmt.Scan(&num)
		if num%5 == 0 && num%11 == 0 {
			fmt.Printf("Divisible %d by 5 and 11", num)
		} else {
			fmt.Printf("Not divisible %d by 5 and 11 ", num)
		}
	}
}

// find the odd number 1 Nth number
func odd1ToN() {
	var num int
	fmt.Print("Enter the number - ")
	fmt.Scan(&num)
	fmt.Print("Odd Numbers  are = ")
	for i := 1; i <= num; i++ {
		if i%2 != 0 {
			fmt.Println(" ", i)
		}
	}
}

// find the power of the number .......
func powerN() {
	var num, exponent int
	pow := 1
	fmt.Print("enter the number - ")
	fmt.Scan(&num)
	fmt.Print("enter the exponent value - ")
	fmt.Scan(&exponent)
	for i := 1; i <= exponent; i++ {
		pow = pow * num
	}
	fmt.Print("power - ", pow)
}

// find the profit and loss
func profit() {
	var costPrice, sellingPrice, ProfitLos int
	fmt.Print("Enter the Cost Price - ")
	fmt.Scan(&costPrice)
	fmt.Print("Enter the Selling Price - ")
	fmt.Scan(&sellingPrice)
	if costPrice > sellingPrice {
		ProfitLos = costPrice - sellingPrice
		fmt.Println("Your Loss Amount - ", ProfitLos)
	} else if sellingPrice > costPrice {
		ProfitLos = sellingPrice - costPrice
		fmt.Println("Your Profit Amount - ", ProfitLos)
	} else {
		fmt.Println("No Profit and No Loss- ", ProfitLos)
	}

}
func main() {
	//divisible()
	//odd1ToN()
	//powerN()
	profit()

}
