package main

import (
	"fmt"
)

func armstrong() {

	var numb, sum, remainder, store int
	fmt.Print("Enter the Number - ")
	fmt.Scan(&numb)
	store = numb   // it storage for number
	for numb > 0 { // it will run till greater than zero 0
		remainder = numb % 10                           // store the remainder after the modulus of by the  10
		sum = (remainder * remainder * remainder) + sum // it is store the sum after the multiplying the remainder and including the sum
		numb = numb / 10                                // this is the store the number after the number dividing by 0
	}
	if store == sum { // here, checking the condition storage is equal to sum
		fmt.Println(sum, "= Armstrong Number ")
	} else {
		fmt.Println(sum, "= Not Armstrong Number ")
	}
}
func main() {
	armstrong()
}
