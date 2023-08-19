package main

import "fmt"

func main() {
	// Declaring the array using the var keyword and
	//accessing of  the array using for loop

	var ar = [5]string{"Chandu", "sonal", "Anjali", "Shruti", "Kajal"}
	//fmt.Println(ar) // this is also print the total array without loop
	Res := ar
	fmt.Print("Total Element -", Res)
	fmt.Println("\n=== Element of the Array ===")
	for i := 0; i < 5; i++ { // Here, i am declaring the for loop
		fmt.Println(ar[i])
	}

}
