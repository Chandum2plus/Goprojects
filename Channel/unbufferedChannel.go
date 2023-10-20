// How to create the unbuffered Channel
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Create a channel
	ch := make(chan string)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		msg := "Namastey Bharat !"
		defer wg.Done()
		fmt.Print("Message sent " + msg)
		ch <- msg
	}()

	go func() {
		defer wg.Done()
		time.Sleep(time.Second)
		rmsg := <-ch
		fmt.Println("message received " + rmsg)
	}()
	wg.Wait()
	
}
