package main

import "fmt"

func main() {
	//Here Creating a slice
	var slic = []int{} // initialize the slice
	slic = append(slic, 10)
	slic = append(slic, 20)
	slic = append(slic, 30)
	slic = append(slic, 40)
	slic = append(slic, 50)

	// converting the slice into the array
	var arr = [5]int{}
	for i, element := range slic {
		arr[i] = element
	}
	fmt.Println("the is Array after converted from slice")
	fmt.Println(arr)
}
