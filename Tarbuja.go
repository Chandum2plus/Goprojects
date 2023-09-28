package main

import "fmt"

// find the even number after its divide
func tarbuja() {
	var num float64
	//var even float64
	for {
		fmt.Print("enter the number - ")
		fmt.Scan(&num)
		var a = int(num)

		if a%2 == 0 {
			var even = int(num)
			if even%2 == 0 {
				fmt.Println("Ha Thik hai")
			}
		} else {
			fmt.Println("E Thik Nahi Hai ")
		}

	}
	
}
func main() {
	tarbuja()
}
