package main

import (
	"fmt"
	"time"
)

func hello() {
	for i := 1; i <= 5; i++ {
		fmt.Println("jay Shree Ram!")
		time.Sleep(time.Second)
	}
}
func main() {
	go hello()

	for i := 1; i <= 5; i++ {
		fmt.Println("Jay Shree Krishna !")
		time.Sleep(time.Second)
	}

}
