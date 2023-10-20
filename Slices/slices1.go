package main

import (
	"fmt"
)

// Calculate the average of the slice
func averageOfSlice() {

	slic := []float64{12, 34, 56, 78, 89}
	fmt.Println("Original Slice -", slic)
	avg := 0.0
	for i := 0; i < len(slic); i++ {
		avg += slic[i] / 5
	}
	fmt.Println("Average of the slice = ", avg)
}

// duplicate element of the slice
func duplicateSlice() {
	sli := []int{23, 34, 32, 23, 45, 66, 77, 66, 54, 34, 90}
	fmt.Println("slice = ", sli)
	dupCount := 0
	fmt.Print("Duplicate Slice =")
	for i := 0; i < len(sli); i++ {
		for j := i + 1; j < len(sli); j++ {
			if sli[i] == sli[j] {
				fmt.Print(" ", sli[i])
				dupCount++
				break
			}

		}
	}
	fmt.Println("\nTotal Duplicate element - ", dupCount)
}

// Arithmetic operation of the slice
func arithmeticSlice() {
	sli1 := []int{12, 22, 34, 55, 67, 78}
	sli2 := []int{11, 23, 44, 43, 22, 12}
	fmt.Println("S1=", sli1)
	fmt.Println("s2=", sli2)
	fmt.Println("Add\tSub\tMult\tDiv\tMod")
	for i := 0; i < 5; i++ {
		fmt.Print("\n", sli1[i]+sli2[i], "\t")
		fmt.Print(sli1[i]-sli2[i], "\t")
		fmt.Print(sli1[i]*sli2[i], "\t")
		fmt.Print(sli1[i]/sli2[i], "\t")
		fmt.Print(sli1[i]%sli2[i], "\t")
	}
	fmt.Println()
}

// Count the positive and negative slice
func countPositiveNegative() {
	s := []int{1, 2, 3, 0, -2, 3, -4}
	fmt.Println("slice =", s)
	count := 0
	fmt.Print("Positive =")
	for i := 0; i < len(s); i++ {
		if s[i] > 0 {
			fmt.Print(" ", s[i])
			count++
		}
	}
	fmt.Println("\nTotal Positive -", count)
	Necount := 0
	fmt.Print("Negative =")
	for _, v := range s {
		if v < 0 {
			fmt.Print(" ", v)
			Necount++
		}
	}
	fmt.Println("\nTotal Negative -", Necount)
}

// largest of the slice
func largestSlice() {
	sl := []int{230, 35, 67, 89, 90}
	fmt.Println("slice =", sl)
	//fmt.Print("Largest slice =")
	large := sl[0]
	small := sl[0]
	sposition := 0
	position := 0
	for i := 0; i < len(sl); i++ {
		if large < sl[i] {
			large = sl[i]
			position = i
		}
	}
	fmt.Println("largest =", large)
	fmt.Println("position =", position)
	for i, v := range sl {
		if small > v {
			small = v
			sposition = i
		}
	}
	fmt.Println("smallest =", small)
	fmt.Println("smallest pos=", sposition)

}
func main() {
	//averageOfSlice()
	//duplicateSlice()
	//arithmeticSlice()
	//countPositiveNegative()
	largestSlice()
}
=================================================
package main

import (
"fmt"
)

// How to find the even and odd number of a slice and sum of the even and odd slice

func evenOdd() {

	slic := []int{2, 4, 5, 6, 7, 7}
	even := 0
	odd := 0
	evensum := 0
	Oddsum := 0
	fmt.Println("Even Number of the slice")
	for i := 0; i < len(slic); i++ {
		if slic[i]%2 == 0 {
			fmt.Println(slic[i])
			even++
			evensum += slic[i]
		}
	}
	fmt.Println("Total Number of the Even =", even)
	fmt.Println("Sum of the Even slice - ", evensum)
	fmt.Println()
	fmt.Println("Number of the Odd slices")
	for i := 0; i < len(slic); i++ {
		if slic[i]%2 != 0 {
			fmt.Println(slic[i])
			odd++
			Oddsum += slic[i]
		}
	}
	fmt.Println("Total Number of the Odd  = ", odd)
	fmt.Println("Sum of the Odd slice - ", Oddsum)
}
func main() {
	evenOdd()
}
===============================================
package main

import (
"fmt"
)

func main() {
	var sli = []int{19, 34, 77, 88, 55, 33, 220}
	fmt.Print(sli)
	var max = sli[0]
	var min = sli[0]
	for i := 0; i < len(sli); i++ {
		for j := i + 1; j < len(sli); j++ {
			if sli[j] > sli[i] {
				max = sli[j]
			} else {
				min = sli[i]
			}
		}
	}
	//for _, ele := range sli {
	//	if ele > max {
	//		max = ele
	//	} else {
	//		min = ele
	//	}
	//}
	fmt.Print("\nLargest Number of the slice - ", max)
	fmt.Print("\nSmallest Number of the slice - ", min)

}
=============================================
package main

import "fmt"

// Reverse the Array
func reverse(input []int) []int {
	var output []int
	for i := len(input) - 1; i >= 0; i-- {
		output = append(output, input[i])
	}
	return output

}
func main() {
	var size int
	fmt.Print("Enter the size of the slice ")
	fmt.Scan(&size)

	var sl = make([]int, size)

	fmt.Print("Enter the element - ")
	for i := 0; i < size; i++ {
		fmt.Scan(&sl[i])
	}
	fmt.Println("Reverse = ", reverse(sl))

}
===========================================
package main

import "fmt"

// Searching the item from the slice
func searching() {
	var size, serchItem, i int
	for {
		fmt.Print("Enter the size of the slice - ")
		fmt.Scan(&size)

		var serchSlic = make([]int, size)
		fmt.Print("Enter the element of the slice - ")
		for i = 0; i < size; i++ {
			fmt.Scan(&serchSlic[i])
		}
		fmt.Print("Enter the Search item - ")
		fmt.Scan(&serchItem)
		flag := 0
		for i = 0; i < size; i++ {
			if serchSlic[i] == serchItem {
				flag = 1
				break
			}
		}
		if flag == 1 {
			fmt.Println("Item is found - ", serchItem)
			fmt.Println("Position of Item - ", i)
		} else {
			fmt.Println(serchItem, " - Item is not found")
		}
	}
}
func main() {
	searching()
}

===========================================================
package main

import (
"bytes"
"fmt"
"reflect"
"sort"
)

// example of the slice
func slice1() {
	slice := []float64{1, 2, 3}
	slice = append(slice, 4, 5, 6)
	fmt.Println("Slices - ", slice)
	sum := 0.0
	for _, v := range slice {
		fmt.Print(" ", v)
		sum += +v
	}
	i := 0
	j := 0
	var temp float64
	for i = 0; i < len(slice); i++ {
		for j = i + 1; j < len(slice); j++ {
			if slice[i] > slice[j] {
				temp = slice[i]
				slice[i] = slice[j]
				slice[j] = temp
			}
		}
	}
	//	fmt.Println("\nGreatest - ", temp)
	Avg := sum / 6
	square := sum * sum
	cube := sum * sum * sum

	fmt.Print("\nsum of the slice - ", sum)
	fmt.Print("\nAverage of the slice -", Avg)
	fmt.Print("\nSquare of the slice - ", square)
	fmt.Print("\nCube of the slice - ", cube)
}

// working of the slice components
func slice2() {
	// Creating an Array
	arr := [6]string{"Chandu", "sonal", "monika", "Anjali", "somi", "Kajal"}
	// Displaying the Array
	fmt.Println("Array -", arr)

	// Creating a slice
	mslice := arr[1:4]

	// Displaying the slice
	fmt.Println("Slice -", mslice)
	fmt.Println("length of the slice - ", len(mslice))

	fmt.Println("Capacity of the slice -", cap(mslice))
}

// Creating a slice using the slice literal
func slice3() {

	// Creating a slice using var keyword
	var mslice1 = []string{"Chandu", "sonal", "somi", "ajali", "saurabh", "kajal"}
	fmt.Println("First Slice - ", mslice1)

	// Creating an Array using the shorthand Declaration

	mlice2 := []int{78, 96, 54, 25, 36, 488, 32, 45}
	fmt.Println("Second Slice - ", mlice2)
}

// How to create the slice from the Array
func slice4() {

	// Creating an Array
	var arr = [5]int{40, 50, 60, 70, 80}

	// Creating the slice from the given Array
	mslice1 := arr[1:2]
	mslice2 := arr[0:]
	mslice3 := arr[:2]
	mslice4 := arr[:]
	fmt.Println("Array - ", arr)
	fmt.Println("Slice1 - ", mslice1)
	fmt.Println("slice2 - ", mslice2)
	fmt.Println("slice3 - ", mslice3)
	fmt.Println("slice4 - ", mslice4)
}

// How to create the slice from the slice
func slice5() {
	var originalSlice = []int{90, 60, 40, 50, 34, 34, 30}
	slice1 := originalSlice[1:5]
	slice2 := originalSlice[0:]
	slice3 := originalSlice[:6]
	slice4 := originalSlice[:]
	slice5 := originalSlice[2:4]

	// Displaying the Result
	fmt.Println("Original Slice - ", originalSlice)
	fmt.Println("New slice1 - ", slice1)
	fmt.Println("New slice2 - ", slice2)
	fmt.Println("New slice3 - ", slice3)
	fmt.Println("New slice4 - ", slice4)
	fmt.Println("New slice5 - ", slice5)
}

// How to create the slice using make function
func slice6() {
	// Creating an array of size 7 and slice this is array till 4 and
	// return the reference of the slice using make function

	var slice1 = make([]int, 4, 7)
	fmt.Printf("slice1 = %v,\nLength1 = %d,\n Capacity1 = %d", slice1, len(slice1), cap(slice1))
	// Creating the another array of size 7 and return the reference of the slice
	// using the make function

	var slice2 = make([]int, 7)
	fmt.Printf("\nSlice2 = %v ,\nLength2 = %d ,\nCapacity2 = %d", slice2, len(slice2), cap(slice2))
}

// How to iterate the over a slice
// first one is using the for loop
func iterateSlice() {

	// Creating the slice using the shorthand declaration
	mslice := []string{"Hi", "this", "is", "Chandu", "Kumar"}

	// iterate over the slice

	for i := 0; i < len(mslice); i++ {
		fmt.Print(" ", mslice[i])
	}
}

// Iterate over the slice using Range for loop
func iterateRange() {

	// Creating the slice using var keyword
	var mslice = []string{"Hi", "this", "is", "Chandu !"}

	// Iteration over a slice using the Range loop
	for index, ele := range mslice {
		fmt.Printf("Index = %d  and Element = %s\n", index+1, ele)
	}
}

// How to iterate over a slice using range in for loop without index value
func iterateWithoutIndexRange() {
	var mslice = []string{"Chandu", "sonal", "saurabh", "somi", "monika", "kajal", "Anjali"}

	// iterating the slice using the range in for loop without index value

	for _, ele := range mslice {
		fmt.Println(ele)
	}
}

// zero value slice
func zeroValueSlice() {
	// Creating a slice with the zero value

	var mslice []string
	fmt.Printf("Length = %d\n", len(mslice))
	fmt.Printf("Capacity = %d\n", cap(mslice))
}

// How to modify the slice
func modiSlice() {
	// Creating an Array
	arr := [6]int{55, 66, 77, 88, 99, 22}
	slic := arr[0:4]
	// before modifying slice and array
	fmt.Println("Array - ", arr)
	fmt.Println("Slice - ", slic)

	slic[0] = 100
	slic[1] = 1000
	slic[2] = 10000
	// After modified the slice
	fmt.Println("\nNew Array - ", arr)
	fmt.Println("New slice - ", slic)
}

// How to check  if slice is nil or not
func sliceNilOrNot() {
	s1 := []int{12, 34, 56}
	var s2 []int
	s3 := []int{12, 54, 56}

	// Here Displaying the value of slice1 and slice2
	fmt.Println("==== Displaying the Value of the both slice =====")
	fmt.Println("Slice 1 = ", s1)
	fmt.Println("Slice 2 = ", s2)
	// You can check like this both slice is equal or not
	fmt.Println("Result = ", reflect.DeepEqual(s3, s1))

	var v, d int
	fmt.Print("slice1 = ", s1)
	fmt.Print("\nslice3 = ", s3)
	fmt.Println()

	for _, v = range s1 {
		fmt.Print(" ", v)

	}
	fmt.Println()
	for _, d = range s3 {
		fmt.Print(" ", d)
	}
	if v == d {
		fmt.Print("\nBoth are equal = ", s3, s1)
	} else {
		fmt.Print("\nBoth are not equal = ", s3, s1)
	}

	// if check like this is then compiler give an error
	// because it is invalid operation s3==s1 the operator is == not defined on []int

	//s3 := []int{13, 23, 34}
	//fmt.Println(s3 == s1)

	// Here checking the slice is nil or not using conditional statement
	/*
		if s1 != nil {
			fmt.Println("Result of slice 1 = ", s1)
		} else {
			fmt.Println("this is nil slice 1 = ", s1)
		}
		if s2 == nil {
			fmt.Println("this is nil slice s2 = ", s2)
		} else {
			fmt.Println("This is not nil slice s2 = ", s2)
		}
	*/
	// Here is also checking the slice is nil or not using the fmt.Println() function
	fmt.Println("Checkin s1 =", s1 == nil)
	fmt.Println("Checkin s2 =", s2 == nil)
}

// How to create the multidimensional slice
// it is just like multidimensional Array expect that slice does not contain the size
func mltDSlice() {
	multiDslice := [][]int{
		{12, 34, 56},
		{45, 67, 55},
		{90, 45, 89},
	}
	fmt.Println()

	fmt.Println(multiDslice)
}

// How to  the sorting a slice using the sort function
func sortSlice() {
	sli := []int{12, 32, 45, 65, 67, 88, 77, 66, 55, 33}
	fmt.Println("Before Sorting = ", sli)
	sort.Ints(sli)
	fmt.Println("After Sorting = ", sli)
	slic := []string{"Chandu", "Anamika", "Neha", "Monika", "Somi"}
	fmt.Print("Before Sorting = ", slic)
	//fmt.Println(slic)
	sort.Strings(slic)

	fmt.Print("\nAfter sorting = ", slic)
	//	fmt.Println(slic)
}

// How to trim the slice  using the bytes trim function
func trimSlice() {

	// Creating and the initializing the byte slice

	slice1 := []byte{'c', 'h', 'i', 'n', 't', 'u', '!', '~', '@'}
	slice2 := []byte{'c', 'h', 'a', 'n', 'd', 'u', '#', '^'}

	// Displaying the slice before the trimming
	fmt.Println("Original slice")
	fmt.Printf("slice1  = %s", slice1)
	fmt.Printf("\nslice2  = %s", slice2)

	// Here, Trimming the slice

	res := bytes.Trim(slice1, "!~@")
	res2 := bytes.Trim(slice2, "#^")

	// Displaying the slice After Trimming
	fmt.Println("\nNew slice ")
	fmt.Printf("Slice1 - %s", res)
	fmt.Printf("\nslice2 - %s", res2)
}
func main () {
	slice1()
	//slice2()
	//slice3()
	//slice4()
	//slice5()
	//slice6()
	//iterateSlice()
	//iterateRange()
	//iterateWithoutIndexRange()
	//zeroValueSlice()
	//modiSlice()
	//sliceNilOrNot()
	//mltDSlice()
	//sortSlice()
	//trimSlice()
}
