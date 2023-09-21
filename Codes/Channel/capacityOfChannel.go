package main

import (
	"fmt"
	"time"
)

// this is the first Goroutine method
func capacity(mchan chan string) {

	mchan <- "Chandu"
	mchan <- "Kajal"
	mchan <- "Anjali"
	mchan <- "Monika"
	close(mchan)
	// capacity  of the channel
	fmt.Println("Capacity = ", cap(mchan))
}

// this is the second Goroutine method
func Number(chnn chan int) {
	//	chnn := make(chan int)
	chnn <- 12
	chnn <- 34
	chnn <- 958
	chnn <- 885
	close(chnn)

	// Here, Iterating over the channel

}

// this is the main function
func main() {
	// Creating a channel using the make function
	mchan := make(chan string, 4) // this is for first Goroutine

	go capacity(mchan) // Calling the function as a Goroutine
	// Iterating Over channel
	for v := range mchan {
		time.Sleep(time.Second)
		fmt.Println(v)
	}
	for { // Checking the condition
		res, ok := <-mchan
		if ok == false {
			fmt.Println("First Channel is Close ", ok)
			break
		}
		fmt.Println("First Channel is Open ", res, ok)
	}

	chnn := make(chan int) // this is for second goroutine

	go Number(chnn) // Calling the second function as a goroutine

	for v := range chnn { // Iterating over second channel
		time.Sleep(time.Second)
		fmt.Println(v)
	}

	for { // Here, checking the condition channel is open or close

		res, ok := <-chnn
		if ok == false {
			fmt.Println("Second Channel is Close ", ok)
			break
		}
		fmt.Println("Second Channel is Open ", res, ok)
	}
	for v := 0; v < 2; v++ { // Here, printing the both Goroutine using the select statement
		select {
		case res1 := <-mchan: // this is for first Goroutine

			fmt.Println("Channel 1 =", res1)
		case res2 := <-chnn: // This is for second Goroutine

			fmt.Println("Channel 2 = ", res2)
		}
	}
}
