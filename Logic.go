package main

import "fmt"

//============================================================================
/* Print the pattern like this -
expected output
 * * * * *
  * * *
   *
  * * *
 * * * * *

*/
func a() {

	for i := 0; i < 3; i++ {
		for sp := 0; sp < i; sp++ {
			fmt.Print(" ")
		}
		for col := 0; col < 5-2*i; col++ {
			fmt.Print(" *")
		}
		fmt.Println()
	}
	for i := 0; i < 2; i++ {
		for sp := 0; sp < 1-i; sp++ {
			fmt.Print(" ")
		}
		for col := 0; col < 2+2*i+1; col++ {
			fmt.Print(" *")
		}
		fmt.Println()
	}

}

//====================================================================================================

/*
Print the pattern like this -
expected output -
* * * * * *
* * * * *
* * * *
* * *
* *
*
*
* *
* * *
* * * *
* * * * *
* * * * * *
*/
func b() {
	for k := 0; k < 6; k++ {
		for v := 0; v < 6-k; v++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
	for i := 0; i <= 5; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}
func res() {
	var name, colname, add string
	fmt.Print("Enter your name -:")
	fmt.Scan(&name)
	fmt.Print("Enter your college name -")
	fmt.Scan(&colname)
	fmt.Print(" Enter your address -")
	fmt.Scan(&add)
	fmt.Println("Your name -", name)
	fmt.Println("Your collage name -", colname)
	fmt.Println("your address -", add)

}
func main() {
	//  a()
	//	b()
	// res()

	//---------------------------------------------------------------

}
