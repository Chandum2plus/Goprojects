package main

import (
	"fmt"
	"strings"
	"time"
)

// Display the Area of the Rectangle with method calling
func area(length, width float64) float64 {
	Ar := length * width
	return Ar
}

// wap the two value using call by value
func swap(a, b int) int {
	var c int
	c = a
	a = b
	b = c
	return c
}

// call by reference
func swap1(a, b *int) int {
	var c int
	c = *a
	*a = *b
	*b = c
	return c

}

// use of the variadic functions
func joinstr(elements ...string) string {
	return strings.Join(elements, "-")
}

// example of variadic function using the slice
func joinStr(elements ...string) string {
	return strings.Join(elements, "-")
}
func main() {
	//fmt.Println(joinStr())
	//element := []string{"Chandu", "kumar", "ritu", "pritam", "sonal"}
	//fmt.Println(joinStr(element...))
	/*
		// zero argument
		fmt.Println(joinstr())
		// multiple arguments
		fmt.Println(joinstr("Hello Chandu"))
		fmt.Println(joinstr("Hello india"))
		fmt.Println(joinstr("Hello World"))

	*/
	/*
		var length, width float64
		fmt.Print("\nEnter the length - ")
		fmt.Scan(&length)
		fmt.Print("Enter the width - ")
		fmt.Scan(&width)
		fmt.Printf("Area of Rectangle = %f\n", area(length, width))

	*/
	/*
		var q int = 10
		var p int = 20
		fmt.Printf("q = %d and p = %d", q, p)
		// Call by value
		swap(q, p)
		fmt.Printf("\nq =%d and p = %d", q, p)

	*/
	/*
		var p int = 10
		var q int = 20
		fmt.Printf("p = %d and q = %d\n", p, q)
		fmt.Println("After swapping")
		// call by reference
		swap1(&p, &q)
		fmt.Printf("p = %d and q = %d", p, q)

	*/
	s := []int{345, 78, 123, 10, 76, 2, 567, 5}
	// sorting the slice
	fmt.Println("sorted slice - ", &s)
	// finding the index
	index := strings.Index("GeeksforGeeks", "gk")
	fmt.Println("index value -", index)
	// finding the time
	fmt.Println("time- ", time.Now().Unix())
}
