package main

import "fmt"

type st struct {
	num int
}

func (s *st) abc() {
	s.num = 90
}
func (s *st) ab() {

	s.abc()

}
func main() {
	var s st // this is the instance of the

	fmt.Println("s.ab =", s.num)
	s.ab()
	fmt.Println("s.abc", s.num)
	//s.abc()

	//s.num = 9
	//fmt.Println(s.num)
}
