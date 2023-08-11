  package main

import "fmt"

func Arr(ar string) {
	r := [100]rune{}
	for i, v := range ar {
		
		r[i] = v
		fmt.Println("rune:=", string(r[i]))
		//fmt.Println(string(v))
	}
}
func main() {
	//ar := "manish"
	var ar string
	fmt.Println("Enter your name")
	fmt.Scanln(&ar)
	Arr(ar)
}
