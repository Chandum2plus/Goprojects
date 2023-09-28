package main

import "fmt"

func Runes() {
	var r = [20]rune{'c', 'h', 'a', 'n', 'd', 'u'}
	for i, v := range r {
		r[i] = v
		fmt.Printf("%c\n", v)
	}
}
func Rune1(name string) {
	for _, v := range name {
		fmt.Printf("%c\n", v)
	}
}
func main() {
	fmt.Println("Namastey Bharat")
	//Runes()
	var n string
	fmt.Print("enter your name -")
	fmt.Scan(&n)
	Rune1(n)

}
