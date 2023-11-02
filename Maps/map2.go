package main

import (
	"fmt"
	"reflect"
)

// Creating and declaring the map using the var keyword

func maps() {

	// Declaring the map using the var keyword
	var emp map[string]interface{}
	fmt.Println(emp)
	fmt.Println(reflect.TypeOf(emp))

	// checking the map empty or not
	if emp == nil {
		fmt.Println("emp Map is nil")
	} else {
		fmt.Println("emp map is not nil")
	}
	emp["Chandu"] = 1
	fmt.Println(emp)
	// initializing the map using the make function
	//emp2 := make(map[string]int, 2)
}
func lenMap() {

	// initialize the map using the shorthand declaration
	scoreRecord := map[string]float64{
		"Chandu":  12.3,
		"Sonal":   23.5,
		"Saurabh": 43.4,
		"Paras":   53,
	}
	if scoreRecord == nil {
		fmt.Println("scoreRecord map is nil ")
	} else {
		fmt.Println("scoreRecord map is not nil ")

	}

	// calculate the length of the map
	getLength := len(scoreRecord)
	fmt.Println("map =", scoreRecord)
	fmt.Println("Length =", getLength)

}

func mapArray() {

	// Initialize the nested map, that means map Array
	mapArr := map[string][]int{
		"Chandu": []int{1, 2, 3, 4, 5},
		"Sonal":  []int{5, 4, 3, 2, 1},
	}
	counter := 0
	for k, v := range mapArr {
		fmt.Printf("%s = %d\n", k, v)
		counter += len(v)
	}
	fmt.Println("counter =", counter)
}
func main() {
	//maps()
	//lenMap()
	mapArray()
}
