package main

import "fmt"

// How to create the buffered channel with the capacity
func main() {

	// Create the channel
	ch := make(chan string, 4) // agr size se jyada data send kroge to deadlock lag jayega
	// agr size jayada hai or data kam hai to koi diikkat nhi hai
	ch <- "chandu"
	ch <- "sonal"
	ch <- "pars"
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
