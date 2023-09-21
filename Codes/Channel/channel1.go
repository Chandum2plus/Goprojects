package main

import "fmt"

// How to create the  channel

func main() {

	// this is the first way of creating the channel using var keyword

	// if create the channel like this using the var keyword then this channel will be null there is no any value
	var mchan chan int // this is null channel
	fmt.Println("value of the Channel ", mchan)
	fmt.Printf("type of the channel %T", mchan)

	// this is the second way of the creating the channel using the shorthand declaration

	// Create the channel like this using the shorthand declaration using the make() function then it channel is initialized
	// there is some value or address
	mchan2 := make(chan string) // this is the initialized channel which has this value or Address 0xc00004e060

	fmt.Println("value of the Channel = ", mchan2)
	fmt.Printf("type of the channel = %T", mchan2)

}
