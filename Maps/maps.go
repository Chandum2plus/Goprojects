package main

import (
	"fmt"
	"sort"
)

// How to create and initialize  the map
func map1() {

	// creating an empty map using var keyword

	var map1 map[int]int
	//fmt.Println(map1)
	// checking the map is empty or not
	if map1 == nil {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	// Creating and initializing the map using the shorthand declaration
	// if you want to create an empty map using the var keyword , so you can't the initialize like this
	// because Golang is already is provided the using the make function. in is procedure value initialization is mendatory
	var map2 = map[int]string{
		1: "dog",
		2: "cat",
		3: "elephant",
		4: "Lion",
		5: "Tiger",
	}
	fmt.Println("map2 is =", map2)

	// initializing the empty map like is abslutelly wrong
	/*
	   var map3=map[int]string{}
	   1: "chandu",
	   2:"sonal"
	   fmt.Println(map3)

	*/
}

// How to create the map using the make function
func map2() {
	// creating the map using the make function
	var mymap = make(map[int]string)
	// this map is empty but you can initialize the value, so you can initialize the
	// value like this is  given below
	mymap[1] = "Chandu"
	mymap[2] = "sonal"
	mymap[3] = "suarabh"

	// if you create  an empty map using the make function, and you want  to initialize the value like this is given below
	// 4: "Anjali", You can't initialize like this because it will give an error
	fmt.Println(mymap)
}

// How to iterate over the map
func iterateMap() {

	// Creating the map using the var keyword
	fmt.Println("Id Name")
	fmt.Println("-------")
	var maps = map[int]string{
		1: "Chandu",
		2: "Sonal",
		3: "Saurabh",
		4: "Manish",
		5: "Somi",
	}
	for id, name := range maps {
		fmt.Println(id, name)
	}
}

// How Add and update the map value
func updatemap() {
	// Creating the map using the shorthand Declaration
	maps := map[int]string{
		1: "cat",
		2: "parrot",
		3: "pekock",
		4: "Crow",
		5: "owl",
	}
	fmt.Println("===== Original map ======")
	fmt.Println(maps)
	// Adding the value
	maps[6] = "Elephant"
	maps[7] = "Lion"
	fmt.Println("After Adding Value")
	fmt.Println(maps)
	// updating the value
	maps[3] = "Mor"
	fmt.Println("After updating the value")
	fmt.Println(maps)

}

// How to Retrieve the value from map using the key
func retrieveValue() {
	// creating the map using the var keyword
	var maps = map[int]string{
		1: "Chandu",
		2: "sonal",
		3: "saurabh",
		4: "Manish",
		5: "Kajal",
	}
	fmt.Println("====Total value of the map ====")
	fmt.Println(maps)
	// Retrieve the value from map using the key
	value1 := maps[1]
	value2 := maps[5]
	fmt.Println("Retrieved value1 = ", value1)
	fmt.Println("Retrieved value2 = ", value2)
	// How to check the  existence of the key in the map
	// like this -
	stdName, ok := maps[5]
	fmt.Println("present or not - ", ok)
	fmt.Println("student name - ", stdName)

	// You can also check the using the blank identifier without value like this -
	_, oks := maps[3]
	fmt.Println("present or not - ", oks)

}

// Creating the map using the interface datatype
func interfacedatatype() {
	m := map[interface{}]interface{}{
		"Name":    "Chandu",
		"Roll":    123,
		"Age":     22.8,
		"College": "Sinha College Aurangabad",
		"Salary":  39484.586,
	}
	fmt.Println("Original map=", m)
	// iteration over a map
	for key, value := range m {
		fmt.Println(key, "- ", value)
	}
}

// How to sort the map
func sortingMap() {
	// key sorting
	unsortedMap := map[string]int{
		"India":  20,
		"Canada": 32,
		"Russia": 55,
		"Japan":  89,
	}
	fmt.Println("Unsorted Map - ", unsortedMap)

	keys := make([]string, 0, len(unsortedMap))
	for k := range unsortedMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, value := range keys {
		fmt.Println(value, unsortedMap[value])
	}

}

// sorting the value of the map
func sortingvalue() {
	maps := map[int]string{
		13: "Sonal",
		22: "Anamika",
		3:  "Monika",
		45: "Somi",
		5:  "Chandu",
		67: "Kajal",
		7:  "Manish",
		81: "Saurabh",
	}
	fmt.Println(maps)
	value := make([]int, 0, len(maps))
	for k := range value {
		value = append(value, k)
	}
	sort.Ints(value)
	for _, k := range value {
		fmt.Println(value, maps[k])
	}
}
func main() {
	//map1()
	//map2()
	//iterateMap()
	//updatemap()
	//	retrieveValue()
	//	interfacedatatype()
	//	sortingMap()
	//	sortingvalue()
	sortingvalue()

}
