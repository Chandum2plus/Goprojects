package main

import (
	"fmt"
	"time"
)

// this is first function named as portal1

func portal1(chanal1 chan string) {
	time.Sleep(time.Second * 2)
	chanal1 <- "welcome to portal1 function !"
	//fmt.Println("welcome to portal1 function !")
}

// this is the second function named as portal2

func portal2(chanal2 chan string) {
	time.Sleep(time.Second * 2)
	chanal2 <- "welcome to portal2 function !"
}

// this is the main function ......
func main() {

	// Creating the channel
	R1 := make(chan string)
	R2 := make(chan string)

	// calling the function1 and function in goroutine
	go portal1(R1)
	go portal2(R2)

	select {
	// Case1 for portal1
	case op1 := <-R1:
		fmt.Print(op1)

		// Case2 for portal2
	case opt2 := <-R2:
		fmt.Println(opt2)

	}

}
