package main

import "fmt"

// Example of the Structure in golang
// Here, Creating the Structure
type student struct {
	name  string
	roll  int
	class string
}

// this is the method name
func struc() {
	a := student{
		name:  "Chandu_kumar",
		roll:  12,
		class: "Sinha College Aurangabad",
	}
	fmt.Println(a)

}
func main() {
	struc()
}
