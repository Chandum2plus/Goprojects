package main

import "fmt"

// Testing the data type of signed integer and unsigned integer
func testingInt() {
	var x uint = 225
	fmt.Println(x, x-3)
	var y int16 = 32767
	fmt.Println(y+2, y-2)

}

func comp() {
	var a complex128 = complex(6, 2)
	var b complex64 = complex(9, 2)
	var sum complex128
	sum = complex(6, 2) + complex(9, 2)
	fmt.Println(sum)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("type of a is %T\n", a)
	fmt.Printf("type of b is %T", b)
}
func comp2() {
	com := complex(10, 11)
	com2 := 13 + 33i
	// Display the Result
	fmt.Println("First complex number is = ", com)
	fmt.Println("Second complex number is = ", com2)
	// get the real number
	realnum := real(com)
	fmt.Println("the real number of the comp =", realnum)
	//get the imaginary number
	imaginary := imag(com2)
	fmt.Println("the imaginary number of the  comp2 =", imaginary)
}
func main() {
	//	testingInt()
	//comp()
	comp2()
}
