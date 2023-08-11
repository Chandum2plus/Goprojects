package main

import (
	"fmt"
)

func alpha(Char rune) {
	//var Char rune
	//fmt.Println("Enter the character - ")
	//fmt.Scanln(&Char)
	//var Char rune
	//fmt.Println("enter the letter ")
	//fmt.Scanln(Char)
	if Char == 'a' || Char == 'e' || Char == 'i' || Char == 'o' || Char == 'u' {
		fmt.Println("this is small vowel")
	} else if Char == 'A' || Char == 'E' || Char == 'I' || Char == 'O' || Char == 'U' {
		fmt.Println("This is capital vowel")
	} else {
		fmt.Println("this is consonant")
	}
	//else if (Char >= 'a' && Char <= 'z') || (Char >= 'A' && Char <= 'Z') {
	//	fmt.Println("This is consonant")
	//}
}
func main() {
	//var Char rune
	//fmt.Println("enter the alpha -")
	//fmt.Scanln(&Char)
	//alpha(Char)
	func() {
		fmt.Println("Hello Anonymous function !")
	}()

}
