package main

import (
	"encoding/json"
	"fmt"
)

type studnt struct {
	Name  string `json:"name"`
	Mob   int    `json:"mob"`
	Roll  int    `json:"roll"`
	Class string `json:"class"`
}

func decodingJson() {

	st := []studnt{
		{"Chandu", 9631812967, 12, "MCA"},
		{"sonal", 9631812987, 13, "BCA"},
		{"saurabh", 9631812909, 14, "MCA"},
		{"manish", 9631812976, 15, "MCA"},
		{"paras", 9631812923, 16, "MCA"},
	}

	data := []byte(`[{"name":Chandu", "mob":9631812967, "roll":12, "class":MCA"},
		{"name":sonal", "mob":9631812987, "roll":13, "class":BCA"},
		{"name":saurabh", "mob":9631812909, "roll":14, "class":MCA"},
		{"name":manish", "mob":9631812976, "roll":15, "class":MCA"},
		{"name":"paras", "mob":9631812923, "roll":16, "class":MCA"}]`)

	err := json.Unmarshal(data, &st)
	if err != nil {
		panic(err)
		//fmt.Println(err)
	}
	for i := range st {
		fmt.Println(st[i])
	}
	fmt.Println("Student =", st)
}
func encAndDec() {
	type student struct {
		Name string
		Roll int
		Age  int
	}
	json_strig := `
{"name":"Chandu","roll":12,"age":23
}`
	st := new(student)
	json.Unmarshal([]byte(json_strig), st)
	fmt.Println("Decoding =", st)

	st1 := new(student)
	st1.Name = "Chandu"
	st1.Age = 23
	st1.Roll = 12

	json_str, err := json.Marshal(st1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Encoding = %s \n", json_str)
}
func main() {
	//decodingJson()
	//encAndDec()

}
