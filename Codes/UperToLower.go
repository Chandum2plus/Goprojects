package main

import (
	"fmt"
	"strings"
)

func UpToLow() {
	var low string
	var up string
	fmt.Printf("If you enter in lower case then result will be in upper case\nIf you enter in upper case then result will be in lower case\n")
	for {
		fmt.Println("Enter in lower case  -")
		fmt.Scanln(&low)
		name := strings.ToUpper(low)
		fmt.Println("Upper case -", name)
		fmt.Println("Enter in upper case -")
		fmt.Scanln(&up)
		res := strings.ToLower(up)
		fmt.Println("lower case -", res)
	}
}

// Example of an Array in golang ....
func array(ar string) {
	r := [100]rune{}
	for i, v := range ar {
		r[i] = v

		fmt.Println("rune = ", string(ar[i]))
	}
	a := strings.ToUpper(ar)
	fmt.Println(a)
}
func main() {
	//UpToLow()
	var arr string
	fmt.Print("Enter Your Name - ")
	fmt.Scan(&arr)
	array(arr)
}
