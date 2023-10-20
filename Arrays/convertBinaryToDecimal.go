package main

import "fmt"

func decimalTobinary(num int) int {

	var Binarray [100]int
	i := 0
	for num > 0 {
		Binarray[i] = num % 2
		num = num / 2
		i++
	}
	for j := i - 1; j >= 0; j-- {
		fmt.Printf("%d", Binarray[i])
	}
	return num
}
func main() {
	var num int
	fmt.Print("Enter the Binary Number - ")
	fmt.Scan(&num)
	fmt.Println(decimalTobinary(num))

}
