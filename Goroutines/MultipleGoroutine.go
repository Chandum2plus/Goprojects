package main

import (
	"fmt"
	"time"
)

// How to create the multiple Goroutine

// This is the first Goroutine
func Aname() {

	// this is the Array of string Datatype
	Arr := [4]string{"Chandu", "suman", "Paras", "Ria"}

	// Iterating on the Array
	for T1 := 0; T1 <= 3; T1++ {
		time.Sleep(150 * time.Millisecond) // set the time

		fmt.Printf("%s\n", Arr[T1])
	}
}

// This is the second Goroutine
func Aid() {

	// This is the Array integer Datatype
	Arr2 := [4]int{300, 301, 302, 303}
	sum := 0
	// Iterating on the Array
	for T2 := 0; T2 <= 3; T2++ {
		time.Sleep(500 * time.Millisecond) // set the time
		fmt.Printf("%d\n", Arr2[T2])
		sum += Arr2[T2]

	}
	time.Sleep(time.Second)
	fmt.Println("sum = ", sum)

}

// This is the main function
func main() {

	fmt.Println("====== Start the main Goroutine ========")

	// Calling the first goroutine
	go Aname() // called as a Goroutine

	// Calling the Second goroutine

	go Aid() // Called as a Goroutine

	time.Sleep(3500 * time.Millisecond)

	//	Aname() // Called as a method
	//	Aid()   // Called as a method

	fmt.Println("====== Main Goroutine End ========")

}
