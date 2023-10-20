package main

import "fmt"

// How to create an array taking the user input
func user() {
	var size, sum, avg int
	fmt.Printf("Enter the size of the Array - ")
	fmt.Scanln(&size)
	var arr = make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Printf("Enter %dth element -: ", i)
		fmt.Scan(&arr[i])
		sum += arr[i]
	}
	avg = sum / size
	fmt.Println("Your Array is =", arr)
	fmt.Println("sum of the Array =", sum)
	fmt.Println("Average of the Array =", avg)

}
func main() {
	user()
}
