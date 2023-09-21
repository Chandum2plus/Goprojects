package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

var (
	partList    = []string{"A", "B", "C", "D"}
	nAssemblies = 3
	wg          sync.WaitGroup
)

func worker(part string) {
	log.Println(part, "Worker begins part")
	time.Sleep(time.Duration(rand.Int63n(1e6)))
	log.Println(part, "Worker completes part")
	wg.Done()
}

// This is the main function
func main() {
	rand.Seed(time.Now().UnixNano())
	for c := 1; c <= nAssemblies; c++ {
		log.Println("Begin Assembly cycle", c)
		wg.Add(len(partList))
		for _, part := range partList {
			go worker(part)
		}
		wg.Wait()
		log.Println("Assemble. cycle", c, "complete")
	}
}
