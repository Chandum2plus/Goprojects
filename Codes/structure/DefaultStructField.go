package main

import "fmt"

// How to use the default value for struct field
type stdnt struct {
	Name    string
	College string
	Address string
	Age     float64
}

// method
func (s *stdnt) info() {
	if s.Name == "" {
		s.Name = "Chandu kumar"
	}
	if s.Address == "" {
		s.Address = "Obra"
	}
	if s.College == "" {
		s.College = "Sinha college"
	}
	if s.Age == 0 {
		s.Age = 22.9
	}
}
func main() {
	s := stdnt{
		Name:    "Paras",
		College: "Sinha college",
		Address: "Madanpur",
		Age:     22.10,
	}
	fmt.Println(s) // display the data
	s.info()       // calling the function first time
	s1 := stdnt{
		Name:    "Chandu kumar",
		College: "Sinha College Aurangabad",
		Address: "Obra",
		Age:     22.9,
	}
	fmt.Println(s1) // printing the Result
	s1.info()       // calling the function second time

}
