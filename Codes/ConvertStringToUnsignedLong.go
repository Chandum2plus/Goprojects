package main

import (
	json2 "encoding/json"
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
)

// Write a program to convert string to unsigned long integer
func stringToUnisignedLong() {
	/*
		convert string type value of "3000" using the parseUint() method and passing
		the value "3000" as its first argument 10 the number base
		as its second argument 32 the bit size as its third argument

	*/
	convetValue, err := strconv.ParseInt("3000", 10, 32)
	// check if any error happened
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// long output of the convertableUintValue
	// variable to the console
	fmt.Println(convetValue)
}

// sorting the array
func sortingArray() {
	fmt.Println("Unsorted array -")
	var arr = []int{7, 8, 9, 6, 5}

	for i := 0; i < len(arr); i++ {
		fmt.Print(" ", arr[i])
	}
	fmt.Println("\nSorted array")
	sort.Ints(arr)
	fmt.Print(arr)
}

// example of rune
func runes(ar string) {
	r := [20]rune{}
	for i, v := range ar {
		r[i] = v
		fmt.Println("rune =", string(r[i]))
	}
}

// second method of rune
func rune2() {
	var name string
	fmt.Print("Enter your name - ")
	fmt.Scan(&name)
	r := [30]rune{}
	for i, v := range name {
		r[i] = v
		fmt.Println("rune = ", string(r[i]))
	}
}

// example of structure in golang
func structure() {
	type Student struct {
		Name string
		Roll int
		Mark float64
	}
	st := Student{}
	fmt.Print("Enter the name of student -")
	fmt.Scan(&st.Name)
	fmt.Print("\nEnter the roll of student -")
	fmt.Scan(&st.Roll)
	fmt.Print("\nEnter the marks of student  -")
	fmt.Scan(&st.Mark)
	fmt.Println("Name - ", st.Name)
	fmt.Println("Roll - ", st.Roll)
	fmt.Println("Marks - ", st.Mark)
}

// example of map in golang
func maps() {

	m := make(map[string]interface{})
	m["name"] = "Chandu kumar"
	m["roll"] = 56
	m["address"] = "aurangabad"
	fmt.Println(m)

}

// example of json in golang
func json() {
	type Salary struct {
		Basic, Hra, Ta float64
	}
	type Employee struct {
		Name          string `json:"name"`
		Add           string `json:"add"`
		Id            int    `json:"id"`
		MonthlySalary []Salary
	}
	data := Employee{
		Name: "Chandu kumar",
		Add:  "obra",
		Id:   123,

		MonthlySalary: []Salary{
			{
				Basic: 233300,
				Hra:   5000,
				Ta:    2000,
			},
		},
	}
	file, _ := json2.MarshalIndent(data, "", " ")
	_ = ioutil.WriteFile("testJoson", file, 0644)
}
func main() {
	json()
	//maps()
	//	structure()
	//rune2()
	/*
		var ar string
		fmt.Println("Enter your name - ")
		fmt.Scanln(&ar)
		runes(ar)

	*/
	//sortingArray()
	//	stringToUnisignedLong()

}
