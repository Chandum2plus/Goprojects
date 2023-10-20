package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getRand() int {
	return rand.Intn(100)
}

// this is Goroutine method
func disp(ch chan interface{}) {

	for i := 1; i <= 10; i++ {
		time.Sleep(time.Second)
		ch <- fmt.Sprintf("Chandu %d", getRand())

	}
	close(ch)
}
func main() {
	mchan := make(chan interface{}, 3)
	// mchan is the special variable that is used to exchange the information among the goroutines

	//close(mchan) // this is very mandatory  closing the channel
	// this is the first way of Receiving the data from channel
	/*
		fmt.Println(<-mchan)
		fmt.Println(<-mchan)
		fmt.Println(<-mchan)

	*/
	// this is the second way of the receiving the data from channel
	go disp(mchan)
	for ch := range mchan {
		//time.Sleep(time.Second)
		fmt.Println(ch)
	}

}
