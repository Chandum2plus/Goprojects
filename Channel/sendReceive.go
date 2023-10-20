package main

import (
	"fmt"
	"sync"
)

func main() {
	mchan := make(chan interface{})
	wg := &sync.WaitGroup{}
	//fmt.Println(<-mchan)
	//mchan <- 5
	wg.Add(2)
	go func(chanl chan interface{}, wg *sync.WaitGroup) {
		fmt.Println("Channel Received - ", <-mchan)
		fmt.Println("Channel Received - ", <-mchan)
		fmt.Println("Channel Received - ", <-mchan)
		for {
			value, isChannelOpen := <-mchan
			if isChannelOpen == false {
				fmt.Println("Channel is Close ", isChannelOpen)
				break
			}
			fmt.Println("Channel is Open ", isChannelOpen, value)

		}
		//fmt.Println(isChannelOpen)
		//fmt.Println(isChannelOpen)
		//fmt.Println(value)
		//fmt.Println(value)
		wg.Done()
	}(mchan, wg)
	go func(chanl chan interface{}, wg *sync.WaitGroup) {
		mchan <- 5
		mchan <- 48
		mchan <- "Chandu kumar"
		close(mchan)
		wg.Done()
	}(mchan, wg)
	wg.Wait()
}
