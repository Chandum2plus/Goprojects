// -----------------------Interfaces with common Method---------------------------
// Two or more interfaces can have one or more common method in list of method sets.
// Here, Structure is a common method between two interfaces Vehicle and Human
package main

import "fmt"

type Vehicle interface {
	structure() []string // this is the common method
	speed() string
}

type Human interface {
	structure() []string // common method
	performance() string
}

// Access the Vehicle
type Car string

// Access the structure() method

func (c Car) structure() []string {
	var parts = []string{"ECU", "Engine", "AirFilter", "wiper", "Wheels", "HeadLight"}
	return parts
}

// Access speed() method

func (c Car) speed() string {
	return "300 km/Hrs"

}

// Access the Human
type Man string

// Access the structure of Human

func (m Man) structure() []string {
	var parts = []string{"Brain", "Heart", "Nose", "Eyelashes", "Leg", "Eye"}
	return parts

}

// Access the performance() method of Human

func (m Man) performance() string {
	return "10 Hrs/day"
}

func main() {
	var c Vehicle // c is instance variable of the Vehicle interface
	c = Car("World Fastest Car BMW")
	var labour Human
	labour = Man("Software Developer")
	fmt.Println("Car's speed:- ", c.speed())
	fmt.Println("Man performance:- ", labour.performance())

	for i, j := range c.structure() {
		fmt.Printf("%-15s <=======> %15s\n", j, labour.structure()[i])

	}
}
