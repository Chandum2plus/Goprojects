package main

import (
	"fmt"
)

func ags() {
	var dd, mm, yyyy int
	var pdd, pmm, pyyyy int
	for {
		fmt.Println("\n========== ENTER THE DATE OF BIRTH ==========")
		fmt.Print("\nEnter the birth of Date - ")
		fmt.Scan(&dd)
		fmt.Print("Enter the birth of Month - ")
		fmt.Scan(&mm)
		fmt.Print("Enter the birth of Year - ")
		fmt.Scan(&yyyy)
		fmt.Println("=========== ENTER THE CURRENT DATE =============")
		fmt.Print("\nEnter the current Date - ")
		fmt.Scan(&pdd)
		fmt.Print("Enter the current Month - ")
		fmt.Scan(&pmm)
		fmt.Print("Enter the current Year - ")
		fmt.Scan(&pyyyy)
		//fmt.Println("Date = ", dd, "Month =", mm, "Year =", yyyy)
		//fmt.Println("Pdate =", pdd, "Pmonth = ", pmm, "Pyear =", pyyyy)

		var d, m, y, days int
		for i := yyyy + 1; i <= pyyyy; i++ {
			if (i%4 == 0 && i%400 == 0) || (i%100 != 0) {
				y = y + 366
			} else {
				y = y + 365
			}
		}
		for i := mm + 1; i <= 12; i++ {
			if i == 1 || i == 3 || i == 5 || i == 7 || i == 8 || i == 12 {
				m = m + 31
			} else if i == 4 || i == 6 || i == 9 || i == 11 {
				m = m + 30
			} else if i == 2 {
				if (yyyy%4 == 0 && yyyy%400 == 0) || (yyyy%100 != 0) {
					m = m + 29
				} else {
					m = m + 28
				}
			}

		}
		if m == 1 || m == 3 || m == 5 || m == 7 || m == 8 || m == 10 || m == 12 {
			for i := dd; i <= 31; i++ {
				d = d + 1
			}
		} else if m == 2 {
			if (yyyy%4 == 0 && yyyy%400 == 0) || (yyyy%100 != 0) {
				for i := dd; i <= 29; i++ {
					d = d + 1
				}
			} else {
				for i := dd; i <= 28; i++ {
					d = d + 1
				}
			}
		}

		for i := 1; i <= pmm; i++ {
			if i == 1 || i == 3 || i == 5 || i == 7 || i == 8 || i == 10 || i == 12 {
				m = m + 31
			} else if i == 4 || i == 6 || i == 9 || i == 11 {
				m = m + 30
			} else if i == 2 {
				if (pyyyy%4 == 0 && pyyyy%400 == 0) || (pyyyy%100 != 0) {
					m = m + 29
				} else {
					m = m + 28
				}
			}
			for i := 1; i <= pdd; i++ {
				d = d + 1
			}

		}
		resD := pdd - dd
		resM := pmm - mm
		resY := pyyyy - yyyy

		days = d + m + y
		fmt.Printf("Days = %d\n", days)
		fmt.Printf("Year : %d Month : %d Day : %d ", resY, resM, resD)
	}
}
func main() {
	ags()

}
