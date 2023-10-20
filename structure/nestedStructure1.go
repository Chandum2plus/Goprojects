package main

import "fmt"

// Here i am declaring the first structure, named as Addresses
type Addresses struct {
	village    string
	city       string
	dist       string
	postalCode int
}

// Here, i am declaring thh second structure, name as Persons
type Persons struct {
	firstName, lastName string
	age                 int
	addres              Addresses // addres is the variable of the Addresses type, Addresses is the user defined data type
}

func show2() {
	res := Persons{
		firstName: "Chandu kumar",
		lastName:  "Pandit",
		age:       22,
		addres:    Addresses{village: "Adma", city: "Aurangabad", dist: "Aurangabad", postalCode: 824121},
	}

	fmt.Println("===== DISPLAYING THE FIRST PERSON ======\n")
	fmt.Println("Name -", res.firstName, res.lastName)
	fmt.Println("Age -", res.age)
	fmt.Println("Village -", res.addres.village)
	fmt.Println("City -", res.addres.city)
	fmt.Println("District -", res.addres.dist)
	fmt.Println("Postal Code -", res.addres.postalCode)

}
func main() {
	show2()
}
