package main

import "fmt"

//	type Math struct {
//		Sum  float64 `json:"sum"`
//		Num1 float64 `json:"num1"`
//		Num2 float64 `json:"num2"`
//	}
//
// func handler() {
//
// }
func main() {
	//fmt.Println("Hello World")
	//fmt.Print("Hello Github !")

	var num int
	fmt.Print("Enter number: ")
	fmt.Scan(&num)
	if num > 0 {
		fmt.Println(num, "is Positive number")
	} else {
		fmt.Println(num, "is Negative number")
	}

}
