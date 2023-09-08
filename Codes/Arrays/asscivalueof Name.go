package main

import (
	"fmt"
)

func names() {
	var name string
	fmt.Print("Enter name - ")
	fmt.Scan(&name)
	sum := name[0]
	for i := 0; i < len(name); i++ {
		fmt.Printf("Ascii value - %c = %d\n", name[i], name[i])
		sum = name[i] + name[i]
	}
	fmt.Println("sum of Ascci -", sum)
}

// Reverse the string
func reverss(str string) string {
	name := []rune(str)

	for i, j := 0, len(name)-1; i < j; i, j = i+1, j-1 {
		name[i], name[j] = name[j], name[i]
	}
	return string(name)
}

// print the even number in Array
func evenInArray() {
	var arr = [5]int{2, 3, 4, 6, 7}
	fmt.Println("array - ", arr)

	//	i := 0
	fmt.Print("\nEven Number - ")
	for i := 0; i < 5; i++ {
		if arr[i]%2 == 0 {
			fmt.Print(" ", arr[i])
		}
	}
	fmt.Print("\nOdd Number - ")
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 != 0 {
			fmt.Print(" ", arr[i])
		}
	}
	//fmt.Print("Even -", even)
	//fmt.Print("Odd -", odd)
}

// Enter the positive and negative number and print in the separate
func negPos() {
	var size int
	fmt.Print("Eneter the size of the Array - ")
	fmt.Scan(&size)
	var arr = make([]int, size)
	var posArr = make([]int, size)
	var negArr = make([]int, size)
	fmt.Print("Enter the element - ")

	for i := 0; i < size; i++ {
		fmt.Scan(&arr[i])
	}
	pos := 0
	neg := 0
	for i := 0; i < size; i++ {
		if arr[i] >= 0 {
			posArr[pos] = arr[i]
			pos++
		} else {
			negArr[neg] = arr[i]
			neg++
		}
	}
	fmt.Println("Postive -", posArr)
	fmt.Println("negtive -", negArr)
}
func main() {
	//names()
	/*
		str := "chandu"
		strRev := reverss(str)
		fmt.Println(str)
		fmt.Println(strRev)
	*/
	//	evenInArray()
	negPos()
}
