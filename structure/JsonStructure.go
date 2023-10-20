package main

import (
	"encoding/json"
	"fmt"
)

// how to access the structure with the json
type StrJson struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	City      string `json:"city"`
}

func Ajson() {
	jsonString := `
{

"fristName":"Chandu",
"lastName":"kumar",
"city":"Patna"

}`
	e := new(StrJson)
	json.Unmarshal([]byte(jsonString), e)
	fmt.Println(e)

	e1 := new(StrJson)
	e1.FirstName = "Paras"
	e1.LastName = "Kumar"
	e1.City = "Madanpur"

	json.Marshal(e1)
	fmt.Printf("%s\n", jsonString)
}

func main() {
	Ajson()

}
