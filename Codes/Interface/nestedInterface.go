// Declares an e1 variable with Employee as its type, then creates a Emp value and assigns it to e1.
package main

import "fmt"

//------------------------Define Type that Satisfies Multiple Interfaces--------------------------------
//Interfaces allows any user-defined type to satisfy multiple interface types at once.
//Using Type Assertion you can get a value of a concrete type back and you can call methods
//on it that are defined on other interface, but aren't part of the interface satisfying.

type Perimeter interface {
	perimeter()
}

type Obj interface {
	numberOfSide()
}

type intType int

// perimeter() method

func (p intType) perimeter() {
	fmt.Println("Perimeter :-", 5*p)
}

// numberOfSide() method

func (p intType) numberOfSide() {
	fmt.Println("IntType has 5 side")
}

func main() {
	var p Perimeter = intType(50)
	p.perimeter()
	var o intType = p.(intType)
	o.numberOfSide()

	//var ob Obj = intType(50)
	//ob.numberOfSide()
	//
	//var pent intType = ob.(intType)
	//pent.perimeter()
}
