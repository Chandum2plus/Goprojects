package main

import "fmt"

func profilose() {
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
func main() {
	profilose()
}
