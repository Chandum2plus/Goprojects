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
