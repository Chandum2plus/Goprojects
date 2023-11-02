package main

import (
	"encoding/json"
	"fmt"
)

/*
Use Field Tags in the Definition of Struct Type
During the definition of a struct type, optional string values may be added to each field declaration.
*/
type students struct {
	Name  string `json:"name"`
	Mob   int    `json:"mob"`
	Roll  int    `json:"roll"`
	Class string `json:"class"`
}

func jsonWithStruct() {

	st := []students{
		{"Chandu", 9631812967, 12, "MCA"},
		{"sonal", 9631812987, 13, "BCA"},
		{"saurabh", 9631812909, 14, "MCA"},
		{"manish", 9631812976, 15, "MCA"},
		{"paras", 9631812923, 16, "MCA"},
	}

	finalJson, err := json.Marshal(st)
	if err != nil {
		panic(err)
	}
	fmt.Println()
	fmt.Printf("%s \n", finalJson)
}
func main() {
	fmt.Println("Welcome to JSON ")
	jsonWithStruct()

}
