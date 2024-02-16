package main

import "fmt"

// print the small alphabet and capital alphabet and also number 0 to 9

func test() {
	//var input rune
	//
	//fmt.Print("Enter value of input rune: ")
	//fmt.Scan(&input)
	//if input >= 'a' && input <= 'z' {
	//	fmt.Println("This is the small alphabetical character")
	//} else if input >= 'A' && input <= 'Z' {
	//	fmt.Println("This is the capital alphabetical character")
	//} else if input >= '0' && input <= '9' {
	//	fmt.Println("This is the numeric value")
	//} else {
	//	fmt.Println("This is the symbol value")
	//}
	//
	//var num int
	//fmt.Print("enter the number - ")
	//fmt.Scan(&num)
	//for i := 1; i <= num; i++ {
	//	fmt.Println(i)
	//}

	var numm int
	i := 2
	fmt.Print("enter the number - ")
	fmt.Scan(&numm)
	for i = 2; i < numm; i++ {
		if numm%i == 0 {
			fmt.Println("Not Prime number is ", i)
		} else {
			fmt.Println(" Prime number is ", i)
		}
	}
}
func main() {

	test()

}
