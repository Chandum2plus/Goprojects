package main

import (
	"fmt"
)

func binary() {
	var num1 int
	fmt.Print("Eneter the First Number - ")
	fmt.Scan(&num1)
	var s int
	for i := 0; i<num1; i++ {
		num1 = num1 / 2
		fmt.Println(s)
		s = num1 % 2
	}

}

func main() {
	binary()

}
