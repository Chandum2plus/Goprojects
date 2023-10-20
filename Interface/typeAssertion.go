package main

import "fmt"

// Implementing the type Assertion
func myfun(a interface{}) {
	val := a.(string)
	fmt.Println("Value - ", val)
}

func main() {
	var val interface {
	} = "Chandu kumar"
	myfun(val)
}
