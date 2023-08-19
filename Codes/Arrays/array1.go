package main

import "fmt"

func main() {
	ar1 := [3]int{10, 29, 30}
	ar2 := [...]int{3, 4, 5, 6, 7, 7, 88, 9}
	ar3 := [5]string{"Chandu", "Sonal", "Shruti", "somi", "monika"}

	// finding the length of  an array using the len()method
	fmt.Println("== PRINT THE LENGTH OF THE ALL ARRAYS ===")
	fmt.Println("Length and Capacity of the Array1 - ", len(ar1), cap(ar1))
	fmt.Println("Length and Capacity of the Array2 - ", len(ar2), cap(ar2))
	fmt.Println("Length and Capacity of the Array3 - ", len(ar3), cap(ar3))
}
