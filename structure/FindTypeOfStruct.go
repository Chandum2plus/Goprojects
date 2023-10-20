package main

import (
	"fmt"
	"reflect"
)

// Find Type of Struct
// The reflect package support to check the underlying type of a struct.

type aa struct {
	length  int
	breadth int
	color   string
}

func findType() {
	r1 := aa{12, 34, "Red"}
	fmt.Println("R1 Type - ", reflect.TypeOf(r1))          //R1 Type -  main.aa
	fmt.Println("R1 Value - ", reflect.ValueOf(r1).Kind()) //R1 Value -  struct

	r2 := aa{
		length:  122,
		breadth: 34,
		color:   "Green",
	}
	fmt.Println("R2 Type - ", reflect.TypeOf(r2))          // 	R2 Type -  main.aa
	fmt.Println("R2 Value - ", reflect.ValueOf(r2).Kind()) // 	 R2 Value -  struct

	r3 := new(aa)
	r3.length = 23
	r3.breadth = 32
	r3.color = "yellow"
	fmt.Println("R3 Type - ", reflect.TypeOf(r3))          //		R3 Type -  *main.aa
	fmt.Println("R3 Value - ", reflect.ValueOf(r3).Kind()) //  R3 Value -  ptr

	r4 := new(aa)
	fmt.Println("R4 Type -", reflect.TypeOf(r4))           //		R4 Type - *main.aa
	fmt.Println("R4 Value - ", reflect.ValueOf(r4).Kind()) // R4 Value -  ptr

	var r5 = &aa{}
	fmt.Println("R5 Type - ", reflect.TypeOf(r5))          // 		R5 Type -  *main.aa
	fmt.Println("R5 Value - ", reflect.ValueOf(r5).Kind()) // 		R5 Value -  ptr

}
func main() {
	findType()

}
