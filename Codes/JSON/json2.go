package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// example of the json using the file creating

func jsonEx() {
	type Student struct {
		Name    string `json:"name"`
		Class   string `json:"class"`
		Roll    int    `json:"roll"`
		College string `json:"college"`
	}
	db := []Student{
		{"Chandu kumar", "BCA", 68, "Sachchidanand Sinha College Auragabad (Bihar)"},
		{"Sona kumar", "BCA", 63, "Sachchidanand Sinha College Auragabad (Bihar)"},
		{"Saurabh kumar", "BCA", 62, "Sachchidanand Sinha College Auragabad (Bihar)"},
		{"Manish kumar", "MCA", 61, "Sachchidanand Sinha College Auragabad (Bihar)"},
		{"Chintu kumar", "BCA", 64, "Sachchidanand Sinha College Auragabad (Bihar)"},
	}
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(db)
	f, err := os.Create("studentDb.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	io.Copy(f, buf)
}
func jsonDec() {
	type Student struct {
		Name    string `json:"name"`
		Class   string `json:"class"`
		Roll    int    `json:"roll"`
		College string `json:"college"`
	}
	db := []Student{}
	f, err := os.Open("studentDb.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	json.NewDecoder(f).Decode(&db)
	fmt.Println("\n", &db)
}
func main() {
	//jsonEx()
	jsonDec()
}
