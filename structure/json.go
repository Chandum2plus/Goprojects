package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// example of the json and json Array
func encodingJsons() {
	type Details struct {
		Name  string `json:"name"`
		Age   int    `json:"age"`
		Email string `json:"email"`
	}
	d := Details{
		Name:  "Chandu kumar",
		Age:   23,
		Email: "chandukumarbca@gmail.com",
	}

	// encoding the json
	data, err := json.MarshalIndent(d, "", "\t")
	if err != nil {
		log.Println(err)
	}
	//
	fmt.Println(string(data))

	// encoding the json Array
	d1 := []Details{
		{Name: "Chandu kumar", Age: 23, Email: "chandukumarbca@gmail.com"},
		{Name: "Rambhu kumar", Age: 23, Email: "rambhukumarbca@gmail.com"},
		{Name: "Sonal kumar", Age: 29, Email: "sonalkumarbca@gmail.com"},
		{Name: "Paras kumar", Age: 23, Email: "paraskumarbca@gmail.com"},
		{Name: "Saurabh kumar", Age: 23, Email: "saurabhkumarbca@gmail.com"},
	}

	data1, err := json.MarshalIndent(d1, "", "\t")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data1))
}
func decodingJsons() {
	type Details struct {
		Name  string
		Age   int
		Email string
	}
	d := Details{
		Name:  "Chandu kumar",
		Age:   23,
		Email: "chandukumarbca@gmail.com",
	}

	// single data decoding
	data := []byte(`{
"Name":"Chnadu kumar",
"Age":23,
"Email":"chandukumarbca@gmail.com"
`)
	err := json.Unmarshal(data, &d)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))

	// Decoding json Array
	//d1 := []Details{
	//	{Name: "Chandu_kumar", Age: 23, Email: "chandukumarbca@gmail.com"},
	//	{Name: "Rambhu_kumar", Age: 23, Email: "rambhukumarbca@gmail.com"},
	//	{Name: "Sonal_kumar", Age: 29, Email: "sonalkumarbca@gmail.com"},
	//	{Name: "Paras_kumar", Age: 23, Email: "paraskumarbca@gmail.com"},
	//	{Name: "Saurabh_kumar", Age: 23, Email: "saurabhkumarbca@gmail.com"},
	//	{Name: "Manish_kumar", Age: 28, Email: "manishkumarbca@gmail.com"},
	//}

	var d1 []Details

	data1 := []byte(`
[
{"Name": "Chandu_kumar", "Age":23, "Email": "chandukumarbca@gmail.com",
"Name": "Rambhu_kumar", "Age":23, "Email": "rambhukumarbca@gmail.com",
"Name": "Sonal_kumar", "Age":29, "Email": "sonalkumarbca@gmail.com",
"Name": "Paras_kumar", "Age":23, "Email": "paraskumarbca@gmail.com",
"Name": "Saurabh_kumar", "Age":23, "Email" :"saurabhkumarbca@gmail.com",
"Name": "Manish_kumar", "Age":28, "Email": "manishkumarbca@gmail.com"}
]`)

	err1 := json.Unmarshal(data1, &d1)

	if err1 != nil {
		fmt.Println(err1)
		//log.Println(err1)
		//panic(err1)
	}
	//fmt.Println(&data1)
	//fmt.Println(string(data1))
	for i := range d1 {
		fmt.Print(d1[i])
		//fmt.Printf("%s", string(data1[i]))
	}
}

func second() {
	// Golang program to illustrate the
	// concept of decoding using JSON

	// declaring a struct
	type Human struct {

		// defining struct variables
		Name    string
		Address string
		Age     int
	}
	// defining a struct instance
	var human1 Human

	// data in JSON format which
	// is to be decoded
	Data := []byte(`{ 
		"Name": "Deeksha", 
		"Address": "Hyderabad", 
		"Age": 21
	}`)

	// decoding human1 struct
	// from json format
	err := json.Unmarshal(Data, &human1)

	if err != nil {

		// if error is not nil
		// print error
		fmt.Println(err)
	}

	// printing details of
	// decoded data
	fmt.Println("Struct is:", human1)
	fmt.Printf("%s lives in %s.\n", human1.Name, human1.Address)

	// unmarshalling a JSON array
	// to array type in Golang

	// defining an array instance
	// of struct type
	var human2 []Human

	// JSON array to be decoded
	// to an array
	Data2 := []byte(` 
	[ 
		{"Name": "Vani", "Address": "Delhi", "Age": 21}, 
		{"Name": "Rashi", "Address": "Noida", "Age": 24}, 
		{"Name": "Rohit", "Address": "Pune", "Age": 25} 
	]`)

	// decoding JSON array to
	// human2 array
	err2 := json.Unmarshal(Data2, &human2)

	if err2 != nil {

		// if error is not nil
		// print error
		fmt.Println(err2)
	}

	// printing decoded array
	// values one by one
	for i := range human2 {

		fmt.Println(human2[i])
	}
}

func main() {
	//encodingJsons() // this is for encoding the data
	decodingJsons()
	//second()
}
