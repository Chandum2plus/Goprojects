package main

import (
	"fmt"
)

func array() {
	arr := [5]int{10, 20, 30, 49, 40}
	max := arr[0]

	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	min := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}
	fmt.Println("max - ", max)
	fmt.Println("Min -", min)
}

// find the second largest number
func secondlargest() {
	var large1, large2 int
	arr := [5]int{23, 12, 45, 67, 87}
	for i := 0; i < len(arr); i++ {
		if large1 < arr[i] {
			large2 = large1
			large1 = arr[i]
		} else if large2 < arr[i] {
			large2 = arr[i]
		}
	}
	fmt.Println("second largest number - ", large2)
}

// store the elements in the array
func storeElements() {
	var size int
	fmt.Print("enter the size of the Array - ")
	fmt.Scan(&size)
	var arr = make([]int, size)
	fmt.Print("enter the elements - ")
	for i := 0; i < size; i++ {

		fmt.Scan(&arr[i])
	}
	fmt.Print("Array - ")
	for i := 0; i < size; i++ {
		fmt.Print(" ", arr[i])
	}
}

/*
2. Write a program in C to read n number of values in an array and display them in reverse order.
Test Data :
Input the number of elements to store in the array :3
Input 3 number of elements in the array :
element - 0 : 2
element - 1 : 5
element - 2 : 7
Expected Output :
The values store into the array are :
2 5 7
The values store into the array in reverse are :
7 5 2
*/
func reverseodr() {
	var size int
	fmt.Print("enter the size of the Array - ")
	fmt.Scan(&size)
	var arr = make([]int, size)
	fmt.Print("enter the element - ")
	for i := 0; i < size; i++ {
		fmt.Scan(&arr[i])

	}
	fmt.Print("Reverse -")
	for i := size - 1; i >= 0; i-- {
		fmt.Print(" ", arr[i])
	}
}

/*
3. Write a program in C to find the sum of all elements of the array.
Test Data :
Input the number of elements to be stored in the array :3
Input 3 elements in the array :
element - 0 : 2
element - 1 : 5
element - 2 : 8
Expected Output :
Sum of all elements stored in the array is : 15
*/
func sumOfAll() {
	var size, sum int
	fmt.Print("enter the size of the Array - ")
	fmt.Scan(&size)
	var arr = make([]int, size)
	fmt.Print("enter the element - ")
	for i := 0; i < size; i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Print("sum of All element - ")
	for i := 0; i < size; i++ {
		sum += arr[i]
	}
	fmt.Print("sum - ", sum)
}

/*
4. Write a program in C to copy the elements of one array into another array.
Test Data :
Input the number of elements to be stored in the array :3
Input 3 elements in the array :
element - 0 : 15
element - 1 : 10
element - 2 : 12
Expected Output :
The elements stored in the first array are :
15 10 12
The elements copied into the second array are :
15 10 12
*/
func copyArray() {
	arr := [5]int{1, 2, 3, 4, 5}
	arr2 := make([]int, len(arr))
	fmt.Print("Original -")
	//i := 0
	for i, v := range arr {
		arr2[i] = v
	}
	fmt.Print("original Arrya - ", arr)
	fmt.Print("copy Array -", arr2)
}

/*
Write a program in C to count the total number of duplicate elements in an array.
Test Data :
Input the number of elements to be stored in the array :3
Input 3 elements in the array :
element - 0 : 5
element - 1 : 1
element - 2 : 1
Expected Output :
Total number of duplicate elements found in the array is : 1
*/
func duplicate() {
	arr := [10]int{11, 33, 5, 11, 5, 3, 4, 9, 9, 2}
	counter := 0
	fmt.Print("duplicates -")
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				fmt.Print(" ", arr[i])
				counter++
				break
			}
		}
	}
	fmt.Print("\nTotal Duplicate - ", counter)
}

// Write a program in C to count the frequency of each element of an array.
func primen() {
	var num int
	fmt.Print("enter the number - ")
	fmt.Scan(&num)
	flag := 0
	for i := 2; i < num; i++ {
		if num%i == 0 {
			flag = 1
			break
		}
	}
	if flag == 1 {
		fmt.Print(num, "- Not a prime number ")
	} else {
		fmt.Print(num, "- is a prime number")
	}
}
func question() {
	for i := 0; i <= 5; i-- {
		fmt.Print(" ", i)
		i++
	}
}
func main() {
	//array()
	//secondlargest()
	//storeElements()
	//reverseodr()
	//sumOfAll()
	//copyArray()
	//duplicate()
	//	primen()
	question()

}
