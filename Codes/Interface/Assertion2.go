package main

import "fmt"

func myfuns(a interface{}) {
	val, ok := a.(float64)
	fmt.Println(val, ok)
}

func main() {
	var val1 interface {
	} = 22.3
	myfuns(val1)

	var val2 interface {
	} = "Chandu kumar"
	myfuns(val2)

}
