package main

import "fmt"

func countdigit() {
	var num int
	count := 0
	sum := 0
	fmt.Print("enter the number - ")
	fmt.Scan(&num)
	for num > 0 {
		num = num / 10
		count = count + 1
	}
	sum += num
	fmt.Printf("total number to digit - %d\n", count)
	fmt.Println("sum of the total number -", sum)
}
func main() {
	countdigit()
}
