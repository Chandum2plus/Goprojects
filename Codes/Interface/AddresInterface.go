//------------------Interface Accepting Address of the Variable---------------------
//The Print() methods accept a receiver pointer. Hence, the interface must also accept a receiver pointer.
//If a method accepts a type value, then the interface must receive a type value; if a method has a pointer
//receiver, then the interface must receive the address of the variable of the respective type.

package main

import "fmt"

type Book struct {
	auther, title string
}
type Magzine struct {
	title string
	issue int
}
type Printer interface {
	print()
}

func (b *Book) Assign(name string, t string) {
	b.auther = name
	b.title = t
}
func (b *Book) print() {
	fmt.Printf("Author of Book :- %s\nTitle of Book :- %s\n", b.auther, b.title)
}
func (m *Magzine) Assign(t string, i int) {
	m.title = t
	m.issue = i
}
func (m *Magzine) print() {
	fmt.Printf("Title of Magazine :- %s\nDate of Issue :- %d", m.title, m.issue)
}
func main() {
	var b Book    // instance of the Book (structure)
	var m Magzine // instance of the Magazine (structure)
	b.Assign("Kalidas", "Ramayana")
	m.Assign("times of Bharat ", 23)
	var p Printer // instance of the Printer (interface)
	fmt.Println("Call the Interface")
	p = &b    // p is store the Address of Book struct
	p.print() // show the Book value via interface
	p = &m    // p is store the Address fo the Magazine Struct
	p.print() // show the value via interface
}
