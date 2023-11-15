package main

import (
	"fmt"
)

//example of the interface using structure

// This is the structure named as input

// Here, I am creating the interface named as student

type student interface {
	PrintName(name string, class string, roll int, course string)
	PrintAddress(village string, post string, pin int, district string)
}
type St int

func (s St) PrintName(name string, class string, roll int, course string) {
	fmt.Println("Student ID -", s)
	fmt.Println("Student Name -", name)
	fmt.Println("Student Class -", class)
	fmt.Println("Student Roll -", roll)
	fmt.Println("Student Course -", course)
	fmt.Println("_______________________________________")
}
func (s St) PrintAddress(village string, post string, pin int, district string) {
	fmt.Println("Village -", village)
	fmt.Println("Post -", post)
	fmt.Println("Pin -", pin)
	fmt.Println("District -", district)

}

// ==================================================================================================================================//
// this is the second example of the interface
type human interface {
	structure() []string
	performance() string
}
type vehicle interface {
	structure() []string
	speed() string
}
type cAr string

func (c cAr) structure() []string {
	var part = []string{"ECU", "Engine", "AirFilter", "Wiper", "GasTask"}
	return part

}
func (c cAr) speed() string {
	return "100km/hr"
}

type mAn string

func (m mAn) performance() string {
	return "10Hr/day"
}

func (m mAn) structure() []string {
	var part = []string{"Brain", "Heart", "Nose", "Eyelashes", "Stomach"}
	return part
}

// =========================================================================================================================
// THIRD EXAMPLE OF THE INTERFACE
/*
//------------------Interface Accepting Address of the Variable---------------------
//The Print() methods accept a receiver pointer. Hence, the interface must also accept a receiver pointer.
//If a method accepts a type value, then the interface must receive a type value; if a method has a pointer receiver,
then the interface must receive the address of the variable of the respective type.
*/

type Books struct {
	BookName   string
	AuthorName string
}
type Magazines struct {
	Title string
	Issue int
}
type priters interface {
	Print()
}

func (b *Books) Assign(n string, a string) {
	b.BookName = n
	b.AuthorName = a
}
func (b *Books) Print() {
	fmt.Println("Author_Name -", b.AuthorName)
	fmt.Println("Book_Name -", b.BookName)
}
func (m *Magazines) Assign(t string, i int) {
	m.Title = t
	m.Issue = i
}
func (m *Magazines) Print() {
	fmt.Println("Magazine_Title -", m.Title)
	fmt.Println("Magazine_Issue -", m.Issue)
}

/*
//-----------------------Empty Interface Type--------------------
//The type interface{} is known as the empty interface, and it is used to accept values of any type. The empty interface doesn't
have any methods that are required to satisfy it, and so every type satisfies it
//The manyType variable is declared to be of the type interface{} and it is able to be assigned values of different types.
//The printType() function takes a parameter of the type interface{}, hence this function can take the values of any valid type.
*/
func emptyInterface(i interface{}) {
	fmt.Println(i)
}
func main() {
	var t student
	t = St(1)
	t.PrintName("Chandu kumar", "BCA", 12, "Computer Application")
	t.PrintAddress("Adma", "Krashawa", 824121, "Aurangabad")

	fmt.Println("______________SECOND OUTPUT_________________")

	var bmw vehicle
	bmw = cAr("World Best Car")
	var dev human
	dev = mAn("Software Developer")
	for i, j := range bmw.structure() {
		fmt.Printf("%-5s <=============>%15s\n", j, dev.structure()[i])
	}

	fmt.Println("______________THIRD OUTPUT__________________")

	var b Books
	var m Magazines
	b.Assign("C,C++", "BalaGuru Shwami")
	m.Assign("Hindu News", 2023)

	var i priters
	i = &b
	i.Print()
	i = &m
	i.Print()

	fmt.Println("____________FOURTH OUTPUT__________________")
	emptyInterface("Chnadu kumar")
	emptyInterface(23)
	emptyInterface(39.98)
	emptyInterface(983940.09873)
	coutry := []string{"Bharat", "Russia", "japan"}
	emptyInterface(coutry)
}
