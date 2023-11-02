package main

import "fmt"

// How to store the Structure into a map
func structToMap() {
	type studentDetails struct {
		Name string
		Roll int
		Age  float64
	}
	stMap := make(map[studentDetails]interface{})
	a := studentDetails{
		Name: "Chandu kumar",
		Roll: 12,
		Age:  23,
	}
	a1 := studentDetails{
		Name: "Rambhu kumar",
		Roll: 13,
		Age:  23,
	}
	stMap[a] = 1
	stMap[a1] = 2
	fmt.Println("\n", stMap)
}
func main() {
	structToMap()
}
