package main

import "fmt"

// main function, from where execution is start
// It must be declared
func main() {
	arr := [5]int{10, 20, 30, 40, 50}
	// Here i have used the length function,
	//with the help it provide the total length of the array
	for i := 0; i < len(arr); i++ { // using the for loop
		fmt.Println(arr[i]) // Displaying the Array
	}
}
