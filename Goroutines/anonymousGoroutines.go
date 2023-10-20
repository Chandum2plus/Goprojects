package main

import (
	"fmt"
	"time"
)

// This is the main function
func main() {

	fmt.Println("welcome to main function !")

	// Creating an Anonymous Goroutine
	go func() {
		fmt.Println("Welcome to M2-Plus !!")
	}()
	time.Sleep(time.Second)
	fmt.Println("Goodbye !! to Main function !")
}
