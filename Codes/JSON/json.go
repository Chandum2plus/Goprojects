package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// example of the JSONs is first
func json1() {
	m := map[string]interface{}{
		"inteValue":   123,
		"floatvalue":  23.45,
		"boolValue":   true,
		"stringValue": "Hello Chandu",
		"objectValue ": map[string]interface{}{
			"arrayValue\n":  []int{23, 4, 5, 6, 7, 8},
			"objectValue\n": []string{"Chandu", "sonal", "Saurabh", "Rambhu"},
		},
	}
	jsonData, err := json.Marshal(m)
	if err != nil {
		fmt.Printf("Coult'nt marshal json date %s\n", err)
		return
	}
	fmt.Printf("Json Data - %s\n", jsonData)
}

// Second example of the JSON
func json2() {
	type Student struct {
		Name      string   `json:"name"`
		College   string   `json:"college"`
		Roll      int      `json:"roll"`
		IsStudent bool     `json:"isStudent"`
		Course    []string `json:"course"`
	}
	jsonString := `{
"name":"Chandu kumar",
"roll":2,
"isStudent":true,
"course":["fundamental","DSA","Go","GitHub"]
}`
	var p Student
	err := json.Unmarshal([]byte(jsonString), &p)
	if err != nil {
		fmt.Printf("Error %s", err)
		return
	}
	fmt.Printf("%+v\n", p)
}

// Third example of the JSON
func json3() {
	type Details struct {
		Name string  `json:"name"`
		Age  float64 `json:"age"`
	}
	d := Details{Name: "Chandu", Age: 22.8}
	jsonbyte, err := json.Marshal(d)
	if err != nil {
		fmt.Printf("Error %s", err)
		return
	}
	fmt.Printf("Data -%s\n", jsonbyte)
}

// example of the fourth json
func json4() {
	type Address struct {
		Vill  string `json:"vill"`
		Post  string `json:"post"`
		PO    string `json:"PO"`
		Pin   int    `json:"pin"`
		Dist  string `json:"dist"`
		State string `json:"state"`
	}
	type Students struct {
		Name  string  `json:"name"`
		Id    int     `json:"id"`
		Roll  int     `json:"roll"`
		Class string  `json:"class"`
		Add   Address `json:"add"`
	}
	st := Students{
		Name:  "Chandu kumar",
		Id:    1,
		Roll:  68,
		Class: "BCA 3rd Year",
		Add: Address{
			Vill:  "Adma",
			Post:  "Krashawan",
			PO:    "Obra",
			Pin:   824121,
			Dist:  "Aurangabad",
			State: "Bihar",
		},
	}
	value, err := json.Marshal(st)
	if err != nil {
		fmt.Printf("Error %s", err)
		return
	}
	//fmt.Printf("value - %s\n", value)
	os.Stdout.Write(value)
}
func main() {
	//json1()
	//json2()
	//json3()
	json4()
}
