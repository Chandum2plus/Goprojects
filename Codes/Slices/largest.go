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
