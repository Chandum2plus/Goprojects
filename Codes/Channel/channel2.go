// How to create the channel send and receive operation
package main

import "fmt"

func sendAndReceive(ch chan int) {
	fmt.Println(234 + <-ch)
}

// main method
func main() {
	fmt.Println("==== Start the Main  Method ====")
	ch := make(chan int)
	go sendAndReceive(ch)
	ch <- 23
	fmt.Println("===== End the Main Method ======")
}
