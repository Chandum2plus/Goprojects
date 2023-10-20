package main

import "fmt"

// Variable storing the Hexadecimal numbers
func variableStore() {

	// Storing the Hexadecimal Number in variable
	x := 0xFF
	y := 0x9C

	// Displaying the values
	fmt.Printf("Type of the x is =  %T\n", x)
	fmt.Printf("Value of the x in Hexadecimal = %X\n", x)
	fmt.Printf("Value of the x in Decimal     = %v\n", x)

	fmt.Printf("Type of the y is = %T\n", y)
	fmt.Printf("Value of the y in Hexadecimal = %X\n", y)
	fmt.Printf("Value of the y in Decimal     = %v\n", y)

}

// How to declaration and initialization of pointers
func initPointer() {

	var a int = 5748 // this is the normal variable
	var s *int       // this is the pointer variable
	s = &a           // It is storing the Address of the A variable

	fmt.Println("value stored in A =", a)               // it's value is 5748
	fmt.Println("Address stored in A =", &a)            // it's Real Address in memory location - 0xc000016098
	fmt.Println("value stored in S   =", s)             // it stored the Address of the A variable - 0xc000016098
	fmt.Println("value stored after Pointing S = ", *s) // it stored the value of the A variable - 5748
	/*
		The default value or zero-value of a pointer is always nil. Or you can say that
		an uninitialized pointer will always have a nil value.
	*/
	var p *int
	fmt.Println("P == ", p)
}

// demonstrate the use of the type interface in pointer variable
func typeInterfaceInPointer() {

	// initializing the normal variable without any type
	var a = 234
	// initializing the pointer variable without any type
	var p = &a

	fmt.Println("value of the A =", a)
	fmt.Println("Address of the A =", &a)
	fmt.Println("Address stored  p =", p)
	fmt.Println("Value stored  P =", *p)
}

// Declaration of the pointer variable using the shortHand declaration
func shortHandPointer() {

	// initializing the normal using the shorthand declaration
	a := 89
	p := &a

	fmt.Println("value of the A =", a)
	fmt.Println("Address of the A =", &a)
	fmt.Println("Address Stored P =", p)
	fmt.Println("value stored P =", *p)
	// It's Output will be like this -
	/*
		value of the A = 89
		Address of the A = 0xc000016098
		Address Stored P = 0xc000016098
		value stored P = 89

	*/
}

// You can also change the value of the pointer or at the
// memory location instead of assigning a new value to the variable.
func valueChange() {

	a := 789 // This is the normal variable

	p := &a // This is the pointer variable which stored the Address of the 'a' variable

	fmt.Println("value of the A=", a)
	fmt.Println("Address of the A=", &a)
	fmt.Println("Address stored  p=", p)
	fmt.Println("Value stored    p=", *p)

	// We can also change the value of the normal variable like this -
	*p = 234
	fmt.Println("After change the value of the A=", a)
	fmt.Println("After change Address of the A =", &a)
	fmt.Println("After change Stored Address P =", p)

}
func main() {
	//variableStore()
	//initPointer()
	//typeInterfaceInPointer()
	//shortHandPointer()
	valueChange()
}
