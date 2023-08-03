package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {

	// finding the max of two numbers

	fmt.Println("max of the two number = ", math.Max(89.9, 78.3))

	// calculate the square root of  a number

	fmt.Println("square root of n number = ", math.Sqrt(67.7))

	// printing the value of pi

	fmt.Println("value of pi = ", math.Pi)

	// Epoch time and milliseconds
	epoch := time.Now().Unix()
	fmt.Println("time = ", epoch)

	// Generating a random integer between 0 to 100
	rand.Seed(epoch)
	fmt.Println(rand.Intn(999))

	//	Finding the Max of two numbers
	//fmt.Println(math.Max(73.15, 92.46))
	//
	//// Calculate the square root of a number
	//fmt.Println(math.Sqrt(225))
	//
	//// Printing the value of `𝜋`
	//fmt.Println(math.Pi)
	//
	//// Epoch time in milliseconds
	//epoch := time.Now().Unix()
	//fmt.Println(epoch)
	//
	//// Generating a random integer between 0 to 100
	//rand.Seed(epoch)
	//fmt.Println(rand.Intn(100))

}
