package main

import "fmt"

func slices() {

	arr1 := [6]int{10, 11, 12, 13, 14, 15}
	myslice := arr1[0:6]

	fmt.Printf("myslice = %v\n", myslice)
	fmt.Printf("length = %d\n", len(myslice))
	fmt.Printf("capacity = %d\n", cap(myslice))
	mslice := []int{2, 3, 4, 5}
	arr := mslice[3]
	for i := 0; i < len(mslice); i++ {
		fmt.Println("array from slice -", mslice[i])
	}
	fmt.Println("array - ", arr)
	fmt.Println("slice -", mslice)
	fmt.Printf("type of array - %T\n", arr)
	fmt.Printf("type of slice - %T\n", mslice)

}

// example of constant ........
func constantTest() {
	const (
		name    = "Chandu"
		id      = 123
		Correct = true
		Salary  = 58000.500
	)
	fmt.Println("Name - ", name)
	fmt.Println("ID - ", id)
	fmt.Println("Correct - ", Correct)
	fmt.Println("Salary - ", Salary)

}

// Declaration of an array
func arrays() {
	var arr = [8]int{45, 63, 89, 90, 95, 96, 97, 99}
	fmt.Println("Total Array element - ", arr)
	fmt.Println(" ---------- Total length of the array --------")
	var max, min int
	for i := 0; i < len(arr); i++ {
		res := arr[i]
		fmt.Printf("Element [%d] = %d\n ", res)
	}
	if arr[7] == max {
		fmt.Println("max - ", max)
	}
	if arr[0] == min {
		fmt.Println("min - ", min)
	}

}

// operation on operators
func operators() {
	a := 4
	b := 5
	fmt.Printf("Bitwise AND %d\n", a&b)
	fmt.Printf("Bitwise OR %d\n", a|b)
	fmt.Printf("Bitwise XOR %d\n", a^b)
}
func main() {
	operators()
	//arrays()
	//constantTest()
	//slices()
	//var p, r, t, si float64
	//fmt.Print("\nenter the amount -")
	//fmt.Scan(&p)
	//fmt.Print("enter the rate -")
	//fmt.Scan(&r)
	//fmt.Print("Enter the time - ")
	//fmt.Scan(&t)
	//si = (p * r * t) / 100
	//fmt.Printf("simple interest = %f", si)
}
