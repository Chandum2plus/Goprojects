package main

import (
	"encoding/json"
	"fmt"
)

// example of the json
func encodingDecodingJSON() {
	type Ck struct {
		Name  string `json:"name"`
		Age   int    `json:"age"`
		Email string `json:"email"`
	}

	// encoding the json
	data := Ck{
		Name:  "Chandu kumar",
		Age:   23,
		Email: "chandukumarbca@gmail.com",
	}

	d, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Encoding =", string(d))

	// Decoding

	var data1 Ck
	d1 := []byte(`{
"name":"Chandu kumar",
"age": 23,
"email":"chandukumarbca@gmail.com"
}`)

	err1 := json.Unmarshal(d1, &data1)
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println("Decoding =", string(d1))

}

// Encoding and decoding of the JSON array

func encodingDecodingJSONArray() {

	type Ck struct { // this  is the structure
		Name  string `json:"name"`
		Age   int    `json:"age"`
		Email string `json:"email"`
	}

	// instance of the Ck for encoding JSON Array
	data := []Ck{
		{Name: "Chandu kumar", Age: 23, Email: "chandukumarbca@gmail.com"},
		{Name: "Rambhu kumar", Age: 23, Email: "rambhukumarbca@gmail.com"},
		{Name: "Sonal kumar", Age: 28, Email: "sonalkumarbca@gmail.com"},
		{Name: "Paras kumar", Age: 23, Email: "paraskumarbca@gmail.com"},
		{Name: "Saurabh kumar", Age: 23, Email: "saurabhkumarbca@gmail.com"},
	}
	// encoding the json Array
	d, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		fmt.Println(err)
	}
	for i := range d {
		fmt.Print(string(d[i]))
	}

	// decoding the JSON Array
	var d1 Ck
	data1 := []byte(`
[
{"nmae": "Chandu kumar", "age": 23, "email": "chandukumarbca@gmail.com"},
{"nmae": "Kajal kumari", "age": 23, "email": "kajalkumarbca@gmail.com"},
{"nmae": "Monika kumari", "age": 23, "email": "monikakumarbca@gmail.com"},
{"nmae": "Somi kumari", "age": 23, "email": "somikumarbca@gmail.com"},
{"nmae": "Reshmi kumari", "age": 23, "email": "reshmikumarbca@gmail.com"},
{"nmae": "Varsa kumari", "age": 23, "email": "varsakumarbca@gmail.com"}

]`)
	err1 := json.Unmarshal(data1, &d1)

	if err != nil {
		fmt.Println(err1)
	}

	for i := range d {
		fmt.Print(string(d[i]))
	}

}
func main() {
	//encodingDecodingJSON()
	encodingDecodingJSONArray()
}
