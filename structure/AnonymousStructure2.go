package main

import (
	"fmt"
)

func main() {
	type student struct {
		int
		string
		float64
	}
	type students struct {
		name  int
		price int
		string
	}
	value1 := students{
		name:   7,
		price:  4,
		string: "Chintuuuuuuuuuu",
	}
	fmt.Println("name -", value1.name)
	fmt.Println("pricr -", value1.price)
	fmt.Println("String -", value1.string)
	value := student{
		int:     78,
		string:  "chandu",
		float64: 80000.50,
	}
	fmt.Println("Roll -", value.int)
	fmt.Println("Name -", value.string)
	fmt.Println("Package -", value.float64)
}
