package main

import "fmt"

func charTest() {
	var data rune

	fmt.Printf("Enter the Number, Character or Special Symbol - ")
	fmt.Scanf("%c", &data)
	fmt.Printf("%c\n", data)
	if (data >= 'a' && data <= 'z') || (data >= 'A' && data <= 'Z') {
		fmt.Printf(" This is the Alphabet %c\n", data)
	} else if data >= '0' && data <= '9' {
		fmt.Printf(" This is the Number %c\n", data)
	} else {
		fmt.Printf("This is the Special Character %c\n", data)
	}

}
func main() {
	charTest()
}
