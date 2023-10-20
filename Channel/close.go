// How to close the channel using the for range loop and close the function
package main

import (
	"fmt"
	"time"
)

// this is the method
func closeChanel(mchanal chan string) {
	for i := 0; i < 4; i++ {
		time.Sleep(time.Second)
		mchanal <- "Chandu kumar"
	}
	close(mchanal)
}

// this is the main function
func main() {

	// Creating the channel
	c := make(chan string)
	// Calling the Goroutine
	go closeChanel(c)
	// When the value of ok is
	// set to true means the
	// channel is open and it
	// can send or receive data
	// When the value of ok is set to
	// false means the channel is closed
	for {
		res, ok := <-c
		if ok == false { // Here, checking the condition
			fmt.Println("Channel is Close ", ok)
			break
		}
		fmt.Println("Channel is Open ", res, ok)
	}
}
