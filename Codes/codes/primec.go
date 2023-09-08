package main

import "fmt"

func main() {
	var num int
	flag := 0
	fmt.Print("enter the  number - ")
	fmt.Scan(&num)
	for i := 2; i < num/2; i++ {
		if num%i == 0 {
			flag++
			break
		}
	}
	if flag == 0 && num != 1 {
		fmt.Println(num, "Prime number")
	} else {
		fmt.Println(num, "Not prime")
	}
}
