package main

import "fmt"

// 1111111
// 1111122
// 1111333
// 1114444
// 1155555
// 1666666
// 7777777
func patt() {
	var num int
	fmt.Print("Enter the Number - ")
	fmt.Scan(&num)
	for row := 1; row <= num; row++ {
		for col := row; col > col-row; col-- {
			fmt.Print("1")
		}
		for row1 := 1; row1 <= num; row++ {
			for col := row1; col <= num; col++ {
				fmt.Print(row1)
			}

		}
		fmt.Println()

	}
}
func main() {
	patt()
}
