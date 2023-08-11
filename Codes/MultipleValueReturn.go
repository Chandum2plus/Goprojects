package main

import "fmt"

func mul_div(num1 int, num2 int) (int, int) {
	return num1 * num2, num1 / num2
}
func main() {

	//mul, div := mul_div(105, 7)
	//// If you are declared div and not used then compiler will give an error , div declared and not used
	//// so must be important to use the both variables mul and div
	//fmt.Printf("Multi 105x7 = %d\nDiv 105/7 = %d", mul, div)
	////fmt.Println("Div - 105 / 7 =", div)

	// using the blank identifier calling the function when you call the using blank identifier then compiler will not give an error
	mul, _ := mul_div(105, 7)
	fmt.Println("105 x 7 =", mul)

}
