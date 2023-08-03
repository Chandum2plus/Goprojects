package main

import (
	"fmt"
	"strconv"
)

func main() {
	//default data type for float value is float64

	//x := 40.5
	//b := int(x)
	//c := x - float64(b) //40.5-40.0
	//fmt.Println(b, c)
	//
	//fmt.Println(int(x), x-float64(int(x)))
	a := 6.5 // this is the floating point datatype, which value is a= 6.5
	b := 5.5 // floating point datatype, value is b= 5.5
	sumflot := a + b
	fmt.Printf("sum of float datatype - %f : %T\n", sumflot, sumflot) // sum of float datatype - 12.000000 : float64

	// after typecasting into the integer type
	c := int(a) // c= 6
	d := int(b) // d = 5

	sumInt := c + d
	fmt.Println("sum of integer part", sumInt) // result = sum of integer part = 11

	// e := float64(c)
	// f := float64(d)
	fmt.Println("sum of float part ", (a-float64(c))+(b-float64(d)))

	//======== CONVERT THE INTEGER TO STRING ==============

	f := 5 // this is integer data type, which value is = 5
	g := 7
	str := string(rune(f))
	str1 := string(rune(g))
	StrResult := str1 + str
	fmt.Printf("sum of string sum = %s : datatype - %T\n", StrResult, StrResult)
	result := strconv.Itoa(f + g) // it will convert into the string datatype
	//result1 := strconv.Itoa(g)
	re := f + g
	fmt.Println("sum of the string = ", re)
	fmt.Printf("type of the before cosnversion - %T\n", f)
	fmt.Printf("after conversion - %T  \n ", result)

	// ====== CONVERT THE FLOAT INTO STRING ==================

}
