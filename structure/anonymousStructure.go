package main

import (
	"fmt"
)

// example of the anonymous function in golang
func main() {
	Element := struct {
		name      string
		branch    string
		language  string
		practicle int
	}{
		name:      "Chandu kumar",
		branch:    "Computer Application",
		language:  "Go language",
		practicle: 498,
	}
	fmt.Println(Element)
	fmt.Println("Name -", Element.name)
	fmt.Println("Branch Name -", Element.branch)
	fmt.Println("Language -", Element.language)
	fmt.Println("Practical -", Element.practicle)
}
