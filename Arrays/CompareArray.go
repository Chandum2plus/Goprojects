package main

import "fmt"

// How to compare the two array in golang

func main() {
	ar1 := [3]int{2, 9, 4}
	ar2 := [...]int{2, 9, 4}
	ar3 := [3]int{2, 8, 5}
	ar4 := [3]int{2, 8, 5}
	fmt.Println("Comparison between the two array")
	fmt.Println("Ar1 == Ar2 :-", ar1 == ar2)
	fmt.Println("Ar2 == Ar3 :-", ar2 == ar3)
	fmt.Println("Ar1 == Ar3 :-", ar1 == ar3)
	fmt.Println("Ar1 == Ar4 :-", ar1 == ar4)
	fmt.Println("Ar2 == Ar4 :-", ar2 == ar4)
	fmt.Println("Ar4 == Ar3 :-", ar4 == ar3)
}
