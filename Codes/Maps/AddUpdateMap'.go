package main

import "fmt"

// How to Add key-value pair in the map and update
func addMap() {
	// Creating the map using the var keyword
	var maps = map[int]string{
		1: "Chandu",
		2: "sonal",
		3: "Sonali",
		4: "Anjali",
	}
	fmt.Println("Original Maps = ", maps)
	// Adding the key-value pair in the same map
	maps[5] = "Rukshar"
	maps[6] = "Kritika"
	maps[7] = "Priyanka"
	fmt.Println("After Adding  = ", maps)

	// Updating the value with help of the key
	maps[5] = "Saurabh"
	maps[6] = "Mainsh"
	fmt.Println("After Updating = ", maps)
}
func main() {
	addMap()
}
