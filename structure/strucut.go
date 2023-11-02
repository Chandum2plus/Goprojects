package main

import (
	"fmt"
)

// example of the structure
func structureSimple() {
	type student struct {
		name   string
		roll   int
		age    int
		email  string
		height float64
	}

	st := student{
		name:   "Chandu_kumar",
		roll:   12,
		age:    23,
		email:  "chandukumarbca@gmail.com",
		height: 5.8,
	}
	fmt.Println("Name -", st.name)
	fmt.Println("Roll -", st.roll)
	fmt.Println("Age -", st.age)
	fmt.Println("Email -", st.email)
	fmt.Println("Height -", st.height)
}

// structure with the zero value
func structureWithZeroValue() {
	type user struct {
		name      string
		roll      int
		age       int
		height    float64
		isAwesome bool
	}
	user1 := user{}
	var user2 user
	fmt.Println(user1)
	fmt.Println(user2)
	user1.name = "Chandu kumar"
	user1.age = 23
	user1.isAwesome = true
	user1.height = 5.8
	fmt.Println(user1)
}

// Example of the structure pointer
func structWithPointer() {
	type vertex struct {
		x int
		y int
	}

	v := vertex{} // instance of the Structure
	v.x = 12
	v.y = 34
	fmt.Println(v.x)
	fmt.Println(v.y)
	p := &v
	p.x = 34
	fmt.Println(p.x)
	(*p).x = 45
	fmt.Println(p.x)

	var (
		v1 = vertex{1, 2}
		v2 = vertex{x: 4}
		v3 = vertex{}
		v4 = &vertex{2, 1}
	)
	fmt.Println("v1 =", v1)
	fmt.Println("v2 =", v2)
	fmt.Println("v3 =", v3)
	fmt.Println("v4 =", v4)

}

/*
The rectangle struct and its fields are not exported to other packages because identifiers are started with an lowercase letter. In Golang, identifiers are exported
to other packages if the name starts with an uppercase letter, otherwise the accessibility will be limited within the package only.

Creating Instances of Struct Types
The var keyword initializes a variable rect. Using dot notation, values are assigned to the struct fields.
*/
func instanceStructure() {
	type rectangle struct {
		length   float64
		breath   float64
		color    string
		geometry struct {
			area      float64
			perimeter float64
		}
	}
	var rect rectangle
	rect.length = 23
	rect.breath = 12
	rect.color = "Red"

	rect.geometry.area = rect.length * rect.breath
	rect.geometry.perimeter = 2 * (rect.length + rect.breath)
	fmt.Println(rect)
	fmt.Println("color =", rect.color)
	fmt.Println("Area =", rect.geometry.area)
	fmt.Println("Perimeter = ", rect.geometry.perimeter)
}

/*
Struct Instantiation Using new Keyword
An instance of a struct can also be created with the new keyword. It is then possible to assign data values to the data fields using dot notation.
*/
func instanceWithNewKeyword() {
	type rectangle struct {

		// creating the structure
		length   int
		breath   int
		color    string
		geometry struct {
			area      int
			perimeter int
		}
	}
	// creating the instance using the new keyword

	var rect = new(rectangle) // rect is pointer the instance of the rectangle
	rect.breath = 23
	rect.length = 46
	rect.color = "Orange"
	rect.geometry.area = rect.breath * rect.length
	rect.geometry.perimeter = 2 * (rect.length + rect.breath)
	fmt.Println(rect)
	fmt.Println("Area =", rect.geometry.area)
	fmt.Println("Perimeter = ", rect.geometry.perimeter)

}

/*
Two instances of the rectangle struct are instantiated, rect1 points to the address of the instantiated struct and rect2 is the name of a struct it represents.

Struct Instantiation Using Pointer Address Operator
Creates an instance of rectangle struct by using a pointer address operator is denoted by & symbol.
*/
func strWithPointer() {

	// creating the structure named as student
	type student struct {
		name  string
		roll  int
		age   float64
		class string
	}
	// creating the instances
	s := &student{}
	s.name = "Chandu"
	s.roll = 22
	s.age = 23
	s.class = "MCA"
	fmt.Println(s)

	s1 := &student{}
	s1.age = 23
	s1.name = "Chintu"
	s1.class = "MCA" // roll skipped
	fmt.Println(s1)
}

/*
Nested Struct Type
Struct can be nested by creating a Struct type using other Struct types as the type for the fields of Struct. Nesting one struct within another can be a useful way to model more complex structures.

*/
func nestedStructure()  {
	type
}
func main() {
	//structureSimple()
	//structureWithZeroValue()
	//structWithPointer()
	//instanceStructure()
	//instanceWithNewKeyword()
	strWithPointer()
}
