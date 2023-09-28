package main

import (
	"fmt"
)

type Alice struct {
	num []int
	str []string
}

func dsa() {
	s := Alice{
		num: []int{12, 34, 56},
		str: []string{"Chandu", "Sonal", "Paras"},
	}
	fmt.Println()
	fmt.Println(s)

}

type S struct {
	b int
}

func second(a *int) {
	*a = 8
	//fmt.Println("Value of second struct", &b)
}
func (s *S) stru() {
	s.b = 58
	fmt.Println("pointer Struct ", &s.b)
}
func main() {
	var r S
	r.b = 57
	r.stru()
	fmt.Println(r.b)
	fmt.Println(r)
	//dsa()
	////	r := S{b: 56}
	////fmt.Println(&r.b)
	////n := Alice{}
	////fmt.Println("in main function", n.num)
	////fmt.Println("in main function", n.str)
	//var a int
	//a = 6
	//second(&a) // call by reference
	//fmt.Println(a)
	//fmt.Println()
}
