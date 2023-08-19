package main

import (
	"fmt"
	"strconv"
	"time"
)

/*
1. Write a program in Go, to store elements in an array and print them.
Test Data :
Input 10 elements in the array :
element - 0 : 1
element - 1 : 1
element - 2 : 2
.......
Expected Output :
Elements in array are: 1 1 2 3 4 5 6 7 8 9
*/
func storeelement() {
	var size int
	fmt.Printf("Enter the size of the Array - ")
	fmt.Scan(&size)
	var arr = make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Printf("Enter %dth Element - ", i)
		fmt.Scan(&arr[i])
	}
	fmt.Print("Array - ", arr)
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
func reverseElement() {
	var size, i int
	fmt.Print("Enter the size of the Array - ")
	fmt.Scan(&size)
	var ar = make([]int, size)
	for i = 0; i < size; i++ {
		fmt.Printf("Enter %dth Element - ", i)
		fmt.Scan(&ar[i])
	}
	for i = 0; i < size; i++ {
		//fmt.Print(" ", ar[i])
	}
	fmt.Println("Array - ")
	for i = size - 1; i >= 0; i-- {
		fmt.Print(" ", ar[i])
	}
	fmt.Println()
	fmt.Print("Reverse = ", ar)

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
func sumAll() {
	var size, sum int
	fmt.Print("Enter the size of the Array - ")
	fmt.Scan(&size)
	var arr = make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Printf("Enter the %dth element - ", i)
		fmt.Scan(&arr[i])
		sum += arr[i]
	}
	fmt.Println("sum =", sum)
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
func copyArr() {
	var size int
	//	var copyar int
	fmt.Print("Enter the size of the element - ")
	fmt.Scan(&size)

	// this is the original array and their size with the help of we can take input from user, and we iterate the array
	var arr = make([]int, size)
	// here, i am which declared,named as copyar with the help of we will copy the array from original array
	var copyar = make([]int, len(arr))
	// it will copy and reverse the array , it means finely will be copied reversed array
	var cpRv = make([]int, len(arr))
	// Here , declaring the loop which iterate the array
	for i := 0; i < size; i++ {
		// Here, taking the input of the element
		fmt.Printf("Enter %dth element - ", i)
		fmt.Scan(&arr[i])
		copyar[i] = arr[i] // this is the copying of the array
	}

	// Here, we declared the loop for reversing the array
	for i := size - 1; i >= 0; i-- {
		//	fmt.Print(arr[i]) // this is the printing the Reverse mod array
		cpRv[i] = arr[i] // and this is copying the array in Reverse mod
		fmt.Print(cpRv[i])
	}
	//fmt.Println("\ncopy reverse =", cpRv) // This is displaying the Result of the copying and Reverse mod
	//	fmt.Println("\nReverse =", arr)      // this is displaying in only reverse array
	fmt.Println("Original Array -", arr) // it is displaying the original array
	fmt.Println("copied =", copyar)      // it is displaying the copied array

}

/*
5. Write a program in C to count the total number of duplicate elements in an array.
Test Data :
Input the number of elements to be stored in the array :3
Input 3 elements in the array :
element - 0 : 5
element - 1 : 1
element - 2 : 1
Expected Output :
Total number of duplicate elements found in the array is : 1
*/
func arrayDuplicate() {
	var size int
	fmt.Print("Enter Size of the Array - ")
	fmt.Scan(&size)
	var arr = make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Printf("Enter %dth element - ", i)
		fmt.Scan(&arr[i])
	}
	var count int
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			if arr[i] == arr[j] {
				count++
				break
			}
		}
	}
	fmt.Println("Array =", arr)
	fmt.Printf("Total Number of duplicate item - %d", count)
}

/*
8. Write a program in C to count the frequency of each element of an array.
Test Data :
Input the number of elements to be stored in the array :3
Input 3 elements in the array :
element - 0 : 25
element - 1 : 12
element - 2 : 43
Expected Output :
The frequency of all elements of an array :
25 occurs 1 times
12 occurs 1 times
43 occurs 1 times
*/
func frequency() {
	var size, counter int
	fmt.Print("Enter size of the Array - ")
	fmt.Scan(&size)
	// if you want to take an array by user input,so you must declare the  using the make function
	// like this -:
	var arr = make([]int, size)
	var freq = make([]int, len(arr))

	// if you want to take an array by the user input, so you can't take using the normal array
	// like this -:
	//var arr = [...]int{}
	//var freq = [...]int{}
	for i := 0; i < size; i++ {
		fmt.Printf("Enter %dth element - ", i)
		fmt.Scan(&arr[i])
		freq[i] = -1
	}
	for i := 0; i < size; i++ {
		counter = 1
		for j := i + 1; j < size; j++ {
			if arr[i] == arr[j] {
				counter++
				freq[j] = 0
			}
		}
		if freq[i] != 0 {
			freq[i] = counter
		}
	}
	fmt.Println("FREQUENCY OF THE ALL ELEMENT AN ARRAY")
	for i := 0; i < size; i++ {
		if freq[i] != 0 {
			fmt.Printf("%d Occures %d time\n", arr[i], freq[i])
		}
	}
}

/*
9. Write a program in C to find the maximum and minimum elements in an array.
Test Data :
Input the number of elements to be stored in the array :3
Input 3 elements in the array :
element - 0 : 45
element - 1 : 25
element - 2 : 21
Expected Output :
Maximum element is : 45
Minimum element is : 21
*/
func maxMin() {
	// this is the user input

	var size, i int
	fmt.Print("Enter size of the Array - ")
	fmt.Scan(&size)
	var arr = make([]int, size)
	for i = 0; i < size; i++ {
		fmt.Printf("Enter %dth Element  - ", i)
		fmt.Scan(&arr[i])
	}
	var min = arr[0]
	var max = arr[0]
	for i = 1; i < size; i++ {
		if arr[i] > max {
			max = arr[i]
		}
		if arr[i] < min {
			min = arr[i]
		}
	}
	fmt.Println("max - ", max)
	fmt.Println("Min - ", min)

	// this is the given input ......
	var arr2 = [5]int{20, 304, 4, 200, 100}
	var min1 = arr2[0]
	var max1 = arr2[0]
	for i := 0; i < len(arr2); i++ {
		if arr2[i] > max1 {
			max1 = arr2[i]
		}
		if arr2[i] < min1 {
			min1 = arr2[i]
		}
	}
	fmt.Println("Max - ", max1)
	fmt.Println("Min - ", min1)

}

/*
10. Write a program in C to separate odd and even integers into separate arrays.
Test Data :
Input the number of elements to be stored in the array :5
Input 5 elements in the array :
element - 0 : 25
element - 1 : 47
element - 2 : 42
element - 3 : 56
element - 4 : 32
Expected Output :
The Even elements are :
42 56 32
The Odd elements are :
25 47
*/
func evenOdd() {
	/*
		var size, i, j, k int
		fmt.Print("Enter the size of the Array - ")
		fmt.Scan(&size)

		var arr = make([]int, size)
		var arr2 = make([]int, size)
		var arr3 = make([]int, size)

		for i = 0; i < size; i++ {
			fmt.Printf("Enter %dth Element - ", i)
			fmt.Scan(&arr[i])
		}
		for i = 1; i < size; i++ {
			if arr[i]%2 == 0 {
				arr2[j] = arr[i]
				j++

			} else {
				arr3[k] = arr[i]
				k++
			}
		}
		for i = 0; i < j; i++ {
			fmt.Println(arr2[i])
		}
		for i = 0; i < k; k++ {
			fmt.Println(arr3[i])
		}
	*/
	var size int
	fmt.Print("Enter size of the Array - ")
	fmt.Scan(&size)

	var arr = make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Printf("Enter %dth Element - ", i)
		fmt.Scan(&arr[i])
	}
	fmt.Println("Even Number -")
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 == 0 {
			fmt.Print(" ", arr[i])
		}
	}
	fmt.Println("\nOdd Number")
	for i := 1; i < len(arr); i++ {
		if arr[i]%2 != 0 {
			fmt.Print(" ", arr[i])
		}
	}

}

/*
11. Write a program in C to sort elements of an array in ascending order.
Test Data :
Input the size of array : 5
Input 5 elements in the array :
element - 0 : 2
element - 1 : 7
element - 2 : 4
element - 3 : 5
element - 4 : 9
Expected Output :
Elements of array in sorted ascending order:
2 4 5 7 9
*/
func sortAscending() {
	var size, temp int
	fmt.Print("Enter the size of the Array - ")
	fmt.Scan(&size)

	var arr = make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Printf("Enter the %dth Element - ", i)
		fmt.Scan(&arr[i])
	}
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			if arr[j] < arr[i] {
				temp = arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}
	for i := 0; i < size; i++ {
		fmt.Print(arr[i])
	}
}

/*
12. Write a program in C to sort the elements of the array in descending order.
Test Data :
Input the size of array : 3
Input 3 elements in the array :
element - 0 : 5
element - 1 : 9
element - 2 : 1
Expected Output :
Elements of the array in sorted descending order:
9 5 1
*/
func sortDescendig() {
	var size, temp int
	fmt.Print("Enter the size of the Array - ")
	fmt.Scan(&size)

	var arr = make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Printf("Enter %dth Element - ", i)
		fmt.Scan(&arr[i])
	}
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			if arr[i] < arr[j] {
				temp = arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}
	for i := 0; i < size; i++ {
		fmt.Print(arr[i])
	}
}

// insertion method in a slice
func insertSlice() {
	var size, i int
	fmt.Print("Enter the size of the Array - ")
	fmt.Scan(&size)
	var arr = make([]int, size)
	for i = 0; i < size; i++ {
		fmt.Printf("Enter %dth element - ", i)
		fmt.Scan(&arr[i])
		fmt.Println(arr[i])
	}
	fmt.Println(arr)
}

/*
15. Write a program in C to delete an element at a desired position from an array.
Test Data :
Input the size of array : 5
Input 5 elements in the array in ascending order:
element - 0 : 1
element - 1 : 2
element - 2 : 3
element - 3 : 4
element - 4 : 5
Input the position where to delete: 3
Expected Output :
The new list is : 1 2 4 5
*/
func deleterArray() {
	var size, pos int
	fmt.Print("Enter the Size of the Array - ")
	fmt.Scan(&size)
	var arr = make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Printf("Enter %dth Element - ", i)
		fmt.Scan(&arr[i])
	}
	fmt.Print("Enter the position where to Delete - ")
	fmt.Scan(&pos)

	for i := 0; i != pos-1; i++ {
		fmt.Println(arr[i])
	}
	for i := 0; i < size; i++ {
		arr[i] = arr[i+1]
		size--
	}

	for i := 0; i < size; i++ {
		//fmt.Print(arr[i])
		fmt.Println(arr)
	}
}

// How to add, subtract days/month/years in golang
func addDaySub() {
	dt := time.Now()
	fmt.Println("time - ", dt) // show the current date and time
	dt1 := dt.AddDate(0, 0, 3)
	fmt.Println("Add three Days - ", dt1) // show date after three days
	dt2 := dt.AddDate(0, 0, -400)
	fmt.Println("Before 400 days - ", dt2)
}

// How to iterate over string
func iteratin(str string) string {
	_, err := strconv.Atoi(str)
	if err != nil {
		return "this is string  -" + str
	} else {
		return "this is Number  -" + str
	}

}

// How to find the duplicate Array in golang
func duplicate() {
	var size int
	fmt.Print("Enter the size of the Array - ")
	fmt.Scan(&size)
	var arr = make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Printf("Eneter %dth Element - ", i)
		fmt.Scan(&arr[i])
	}
	//var arr = [10]int{10, 20, 30, 20, 10, 40, 88, 22, 62, 22}
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				fmt.Println(arr[j])
			}
		}
	}
}

// How to Add the two Array
func addArray() {
	var size int
	fmt.Print("Enter the size of the Array - ")
	fmt.Scan(&size)

	// Here,declaring the Array -
	var arr1 = make([]int, size)
	var arr2 = make([]int, size)
	var AddAr = make([]int, size)
	fmt.Print("Enter First Array of Element  - ")

	for i := 0; i < size; i++ {
		fmt.Scan(&arr1[i])
	}
	fmt.Print("Enter Second Array of Element - ")

	for i := 0; i < size; i++ {
		fmt.Scan(&arr2[i])
	}
	fmt.Print("Addtion of the Array - ")
	for i := 0; i < size; i++ {
		AddAr[i] = arr1[i] + arr2[i]
		fmt.Print(AddAr[i], " ")
	}
	fmt.Println()
}

// Arithmetic operation Between the Two Array ....
func arithmeticArray() {

	// This is the given input -
	/*
		var arr = [6]int{10, 29, 30, 40, 50, 67}
		var arr2 = [6]int{47, 58, 34, 23, 21, 43}
		fmt.Println("Add\tSub\tMul\tDiv\tMod\t")
		for i := 0; i < 6; i++ {
			fmt.Print("\n", arr[i]+arr2[i], "\t")
			fmt.Print(arr[i]-arr2[i], "\t")
			fmt.Print(arr[i]*arr2[i], "\t")
			fmt.Print(arr[i]/arr2[i], "\t")
			fmt.Print(arr[i]%arr2[i], "\t")
		}
		fmt.Println()

	*/

	// Arithmetic operation using the user input -
	var size int
	fmt.Print("Enter the Size of the array - ")
	fmt.Scan(&size)

	var arr1 = make([]int, size)
	var arr3 = make([]int, size)

	fmt.Print("Enter the First Array Element - ")
	for i := 0; i < size; i++ {
		fmt.Scan(&arr1[i])
	}
	fmt.Print("\nEnter the Second Array Element - ")

	for j := 0; j < size; j++ {
		fmt.Scan(&arr3[j])
	}
	fmt.Println("Add\tSub\tMulti\tDiv\tMod\t")
	for i := 0; i < size; i++ {
		fmt.Print("\n", arr1[i]+arr3[i], "\t")
		fmt.Print(arr1[i]-arr3[i], "\t")
		fmt.Print(arr1[i]*arr3[i], "\t")
		fmt.Print(arr1[i]/arr3[i], "\t")
		fmt.Print(arr1[i]%arr3[i], "\t")
	}
	fmt.Println()
}

// find the Even or odd number in Array
func evod() {
	var arr = [5]int{0, 2, 3, 4, 6}
	fmt.Print("Even Number -: ")
	for i := 0; i < 5; i++ {
		if arr[i]%2 == 0 {
			fmt.Print(arr[i])
		}
	}
	fmt.Print("\nOdd Number -: ")
	for i := 0; i < 5; i++ {
		if arr[i]%2 != 0 {
			fmt.Print(arr[i])
		}
	}
}
func main() {
	//storeelement()
	//reverseElement()
	//sumAll()
	//copyArr()
	//arrayDuplicate()
	//frequency()
	//maxMin()
	//evenOdd()
	//sortAscending()
	//sortDescendig()
	//insertSlice()
	//deleterArray()
	//	addDaySub()
	//var str string
	//fmt.Print("Enter the string/number - ")
	//fmt.Scan(&str)
	//fmt.Println(iteratin(str))
	//duplicate()
	//addArray()
	//arithmeticArray()
	//evod()
}
