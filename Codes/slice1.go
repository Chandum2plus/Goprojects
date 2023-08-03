package main

import "fmt"

func main() {
	var s = []int{12, 36, 78, 96, 25}
	fmt.Print(s)
	m := s[2:3]
	fmt.Println(m)
}
