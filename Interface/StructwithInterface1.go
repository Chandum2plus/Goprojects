package main

import "fmt"

// =================== Implement the interface with the  structure =================

// create the interface
type Vehicles interface {
	mileage() float64
	speed() float64
}
type BMW struct {
	fuel         float64
	distance     float64
	averageSpeed float64
}

type b BMW

func (b BMW) mileage() float64 {

	return b.mileage()
}
func (b BMW) speed() float64 {
	return b.speed()
}
func main() {
	var v Vehicles
	var b BMW
	b.fuel = 122
	b.distance = 2445
	b.averageSpeed = 78
	//fmt.Println("fuel =", b.fuel)
	//fmt.Println("distance =", b.distance)
	//fmt.Println("average speed =", b.averageSpeed)

	fmt.Println("fuel =", b.fuel)
	fmt.Println("distance =", b.distance)
	fmt.Println("average =", b.averageSpeed)
	v.speed()
	v.mileage()

}
