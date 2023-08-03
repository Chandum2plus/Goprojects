package main

import "fmt"

/*
we can declare an array that contains pointers as its elements. every element of this array is a pointer variable that can hold
address of any variable of appropriates type. the syntax of declaring an array of pointers is similar to that of declaring
arrays except that an asterisk of before the datatype name
*/

func main() {

	// Here we are declaring the variable with value
	a := 5
	b := 10
	c := 15

	// Here we are declaring the array of pointer. which is empty array here is no element only defining the size
	pa := [3]*int{}
	pa[0] = &a
	pa[1] = &b
	pa[2] = &c

	// Here use the for loop for iterate the array lengths
	fmt.Println("========== This is the Output =======")
	for i := 0; i < 3; i++ {
		fmt.Printf("pa [%d] = %p\t ", i, pa[i])
		fmt.Printf("*pa [%d] = %d \n", i, *pa[i])
	}
	/*
	   Here pa is declared as an array of pointers. Every element of this array is a pointer to an integer.
	   pa[i] gibes the value of the ith element of pa which is an address of any int variable and *pa[i] gives the value
	   of that int variable.the array of pointers can also contain addresses of elements of another array.
	*/
}
