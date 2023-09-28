package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// producer
func sender(ch chan string) {
	defer wg.Done()
	msg := "hello bharat"
	msg2 := "Hello Chandu"
	ch <- msg
	ch <- msg2
	fmt.Println("message sent " + msg)
	fmt.Println("message2 sent " + msg2)

	//close(ch)

}

// consumer
func receiver(ch1 chan string) {
	wg.Done()

	rmsg := <-ch1
	rmsg2 := <-ch1
	fmt.Println("message Received " + rmsg)
	fmt.Println("message2 Received " + rmsg2)

	//close(ch1)
}
func simple(ch3 chan int) {
	defer wg.Done()
	ch3 <- 48
	fmt.Println(ch3)
}
func main() {
	wg.Add(2)
	ch2 := make(chan string)

	go sender(ch2)
	go receiver(ch2)
	ch3 := make(chan int)
	go simple(ch3)
	fmt.Println(<-ch3)
	wg.Wait()

}
