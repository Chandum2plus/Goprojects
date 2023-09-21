package main

import (
	"fmt"
	"time"
)

func Chanl(chanl chan string) {

	chanl <- "Chandu"
	chanl <- "sonal"
	chanl <- "Paras"
	chanl <- "saurabh"
	close(chanl)
	fmt.Println("length -", len(chanl)) // its length shows 3 because its start the indexing from zero 0

}

// main function
func main() {

	// Creating a channel
	ch := make(chan string, 4)
	go Chanl(ch)
	for v := range ch {
		time.Sleep(time.Second)
		fmt.Println(v)
	}

	for {
		res, ok := <-ch
		if ok == false {
			fmt.Println("channel is close = ", ok)
			break
		}
		fmt.Println("channel is open =", res, ok)
	}

}
