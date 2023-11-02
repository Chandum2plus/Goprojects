package main

import "fmt"

// Example Structure with map
func structWithMap() {
	type student struct {
		name  string
		roll  int
		class string
	}
	a1 := student{
		name:  "Chandu kumar",
		roll:  12,
		class: "MCA",
	}
	a2 := student{
		name:  "Sonal",
		roll:  13,
		class: "BCA",
	}
	var mp map[student]int
	if mp == nil {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
	sm := map[student]int{a1: 1, a2: 2}
	fmt.Println(sm)
	m := make(map[int]string, 10)
	m[1] = "kutta"
	fmt.Println(m)
}
func main() {
	structWithMap()
}
