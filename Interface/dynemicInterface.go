package main

import "fmt"

// Creating an interface
type tank interface {
	Tarea() float64
	Volume() float64
}

func main() {
	var t tank
	fmt.Println("value of the Tank - ", t)
	fmt.Printf("type of the volume %T", t)

}
