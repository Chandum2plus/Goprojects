package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println()
	ch := make(chan string, 2)
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		msg1 := "Hello chandu"
		msg2 := "Nice to meet you"
		defer wg.Done()
		fmt.Println("message send " + msg1 + "and " + msg2)
		ch <- msg1
		ch <- msg2
	}()
	go func() {
		defer wg.Done() // delay the execution
		rmsg := <-ch
		rmsg2 := <-ch
		fmt.Println("message Received " + rmsg + "and " + rmsg2)
	}()
	wg.Wait()
}
