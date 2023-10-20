package main

import "fmt"

type Address struct {
	name, village, dist, state string
	pincode                    int
}

func (a Address) address() {
	fmt.Println("\n====== RESULT ======")
	fmt.Println("Name - ", a.name)
	fmt.Println("Village - ", a.village)
	fmt.Println("District - ", a.dist)
	fmt.Println("State - ", a.state)
	fmt.Println("Pin code - ", a.pincode)
}
func main() {
	//res := Address{           // This is the given value , it means first way
	//	name:    "Chandu",
	//	village: "Adma",
	//	dist:    "Aurangabad",
	//	state:   "Bihar",
	//	pincode: 824121,
	//}
	//res.address()

	// this is the second way using the user input
	res := Address{}
	fmt.Print("Enter the Name - ")
	fmt.Scan(&res.name)
	fmt.Print("Enter the Village - ")
	fmt.Scan(&res.village)
	fmt.Print("Enter the District - ")
	fmt.Scan(&res.dist)
	fmt.Print("Enter the State - ")
	fmt.Scan(&res.state)
	fmt.Print("Enter the Pin code - ")
	fmt.Scan(&res.pincode)
	res.address()
}
