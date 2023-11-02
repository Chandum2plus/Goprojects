package main

import (
	"fmt"
	"strings"
)

// Concept of the Variadic Function
func variadic(str ...string) string {
	return strings.Join(str, "-")
}
func main() {

	// calling the function without any argument
	fmt.Println(variadic())

	// calling the function with argument
	fmt.Println(variadic("Chandu", "Kumar", "Paras"))
}
