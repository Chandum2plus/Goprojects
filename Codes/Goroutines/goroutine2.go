package main

import (
	"fmt"
	"time"
)

func name(str string) {
	for i := 1; i <= 5; i++ {
		time.Sleep(time.Second)
		fmt.Println(str)
	}
}
func main() {
	go name("Chandu")
	name(" kumar")

}
