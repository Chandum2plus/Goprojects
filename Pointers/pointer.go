package main

import "fmt"

// example of the pointers single,double and multilevel pointers ..........

func pointer() {
	var a int = 12
	var p *int = &a
	var p2 **int = &p
	fmt.Println("value of a = ", a)
	fmt.Println("Address of a = ", &p)
	fmt.Println("Address of p = ", *p2)
	fmt.Println("Address of the P = ", &p2)
	fmt.Println("Value of the A = ", *p)
	
}
func main() {
	pointer()
}
