package main

import (
	"fmt"
	"sort"
)

// How to create a map using the instance variable

func mp1() {
	// creating a map using the instance variable which have string as key-type and value as a int type
	var emp = map[string]int{"Chandu": 22, "sonla": 29}
	fmt.Println(emp) // print the Result
}

// Creating an empty map
func mp2() {
	var emp = map[string]int{}
	fmt.Println(emp)
	fmt.Printf("type %T\n", emp)
	// Result will be like this
	//	map[]
	//	type map[string]int

}

// How to declare the map using the make function
func mp3() {
	// map declaration using the make function
	var emp = make(map[string]int)
	emp["Chandu"] = 10
	emp["Sonal"] = 20
	fmt.Println(emp)

	// map declaration using the shorthand declaration
	emplist := make(map[string]int)
	emplist["Paras"] = 30
	emplist["Saurbh"] = 40
	emplist["Manis"] = 50
	fmt.Println(emplist)

	v := make(map[string]int)
	fmt.Println(v)
	fmt.Println("Length of Emp = ", len(emp))
	fmt.Println("Length of Emplist = ", len(emplist))
	fmt.Println("Length of V = ", len(v))

}

// Accessing the item from map
func mp4() {
	var emp = map[string]int{"Chandu": 10, "Raju": 30}
	fmt.Println(emp["Chandu"]) // Accessing the item using the key
}

// iterate over the map
func iterate() {
	var mp = map[string]int{"Chandu": 12, "Sonal": 23, "Paras": 34, "Saurabh": 90}
	fmt.Println("key      value")
	for k, v := range mp {
		fmt.Println(k, "->", v)
	}
}

// truncate the value form map
func truncate() {
	var mp = map[string]int{"Chandu": 12, "Sonal": 23, "Paras": 34, "Saurabh": 90}
	fmt.Println("key      value")
	for k := range mp {
		delete(mp, k)
	}
}

// sort map
func sorted() {
	mp := map[string]int{"chandu": 10, "Sonal": 20, "Ram": 100, "Sita": 89}

	key := make([]string, 0, len(mp)) // this is the slice

	for value := range mp {
		key = append(key, value)
	}
	sort.Strings(key)
	for _, k := range key {
		fmt.Println(k, mp[k])
	}
}
func main() {
	//mp1() // calling the method
	//mp2() // calling the method
	//mp3() // calling the function
	//mp4()
	//iterate()
	//truncate()
	sorted()
}
