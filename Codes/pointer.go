package main

import "fmt"

// write a program to print the value and address of the element of the array
func Arr() {
	arr := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("===== Print the Array with value and Address =====")
	for i := 0; i < len(arr); i++ {
		fmt.Printf("value of array - [%d] = %d \t", i, arr[i])
		fmt.Printf("Address of array - [%d] = %p \n", i, &arr[i])
	}
}

// write a program to print the value and address of element of an array using pointer
func Arr1() {
	arr := []int{10, 20, 30, 40, 50}
	fmt.Println("========= Print the Array value with Address using Pointer =========")
	fmt.Println("The Array is - ", arr)
	for i := 0; i < len(arr); i++ {
		fmt.Printf("Value of array - [%d] = %d\t", i, arr[i])
		fmt.Printf("Address of array - [%d] = %p\n", i, &arr[i])
	}

}

// write a program to print the call by value
func value(x, y int) {

	x++
	y++
	fmt.Printf("inside the function x = %d,y =%d\n", x, y)
}

// write a program to print the call by reference .
func callByReference(p *int, q *int) {
	(*p)++
	(*q)++
	fmt.Printf("Inside the function *p = %d and *q = %d\n", *p, *q)

}

// write a program to show a function that returns pointers
func poiter(p, n *int) *int {
	p = p + n
	return p
}
func main() {
	// Arr()
	//Arr1()
	//------------------------------------------//
	// Here passing the call by value .......
	/*
		a := 5
		b := 8
		fmt.Printf("Before calling the function, a = %d and b = %d\n", a, b)
		fmt.Printf("After calling the function, a = %d and b = %d\n", a, b)

	*/
	//	value(a, b)
	//------------------------------- ----------------------------------

	// Here passing the call by reference
	/*
		c := 6
		d := 9
		fmt.Printf("Before calling the function , a = %d and b = %d \n", c, d)
		callByReference(&c, &d)
		fmt.Printf("After calling the fucntion, a = %d and b = %d \n", c, d)

	*/

	// Here passing the array which is returns the function pointers
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	n := 5

	var ptr *int
	ptr = poiter(arr, &n)
	fmt.Printf("Array = %p,ptr = %p, *ptr = %d\n", arr, ptr, *ptr)

}
