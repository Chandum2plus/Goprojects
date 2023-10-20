package main

import "fmt"

func main() {
	name := "Chandu"
	fmt.Print("Total Length = ", len(name))

	// concept of the unidirectional channel
	mcha := make(<-chan string) //Only for Receive

	mchn := make(chan<- string) // Only for Send

	fmt.Printf("\n Only for Receive - %T", mcha)
	fmt.Printf("\n Only for Send    - %T", mchn)

}
