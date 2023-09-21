package main

import "fmt"

// find the sum of the individual digits of a positive number
func individual() {
	var num int
	fmt.Print("enter the individual number - ")
	fmt.Scan(&num)
	sum := 0
	count := 0
	for num > 0 {
		sum += num % 10
		num = num / 10
		count++
	}
	fmt.Println("sum = ", sum)
	fmt.Println("count =", count)

}
func main() {
	individual()
}
