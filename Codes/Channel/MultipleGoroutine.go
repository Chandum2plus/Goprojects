package main

import (
	"fmt"
	"time"
)

// this is main function
func main() {

	// Creating a channel using the make function

	mchan := make(chan string)

	// this is the anonymous function
	go func() {

		// Here sending the element to channel
		mchan <- "Chandu kumar"
		mchan <- "Paras"
		mchan <- "Sonal"
		mchan <- "Saurabh"
		close(mchan)
	}()

	// Iterating the channel using the for loop

	for Res := range mchan {
		time.Sleep(time.Second)
		fmt.Println(Res)
	}

	// Here, checking the condition
	for {
		res, ok := <-mchan
		if ok == false {
			fmt.Println("Channel is Close ", ok)
			break
		}
		fmt.Println("Channel is Open", res, ok)
	}
}
