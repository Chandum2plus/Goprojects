package main

import "fmt"

// =============================== Example of the structure with interface =======================
type Studnt struct {
	name string
	roll int
	age  float64
}
type Printers interface {
	print()
}

// type pr Printers
func (s *Studnt) Assign(n string, r int, a float64) {
	s.name = n
	s.roll = r
	s.age = a
}
func (s *Studnt) print() {
	fmt.Println("Name -", s.name)
	fmt.Println("Roll -", s.roll)
	fmt.Println("Age  -", s.age)
}
func main() {
	var s Studnt
	s.Assign("Chandu kumar", 12, 22.9)
	var pr Printers // pr is instance of Printers
	pr = &s         // and pr is stored  the address of the s it means student structure
	pr.print()      // Here, calling the print() function

}
