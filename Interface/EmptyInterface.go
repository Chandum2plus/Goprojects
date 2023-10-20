//-----------------------Empty Interface Type--------------------
//The type interface{} is known as the empty interface, and it is used to accept values of any type. The empty interface doesn't
//have any methods that are required to satisfy it, and so every type satisfies it

//The manyType variable is declared to be of the type interface{} and it is able to be assigned values of different types.
//The printType() function takes a parameter of the type interface{}, hence this function can take the values of any valid type.

package main

import "fmt"

func printType(i interface{}) {
	fmt.Println(i)
}
func main() {
	var manyType interface{}
	manyType = "Chandu kumar"
	fmt.Printf("%s -> type %T\n", manyType, manyType)

	manyType = 12
	fmt.Printf("%d -> Type %T\n", manyType, manyType)

	manyType = 239.48
	fmt.Printf("%f -> type %T\n", manyType, manyType)
	// 1 slice
	// 2 map
	// 3 array
	//a := "chandu_kumar"
	printType("Chandu kumar")
	var countries = []string{"Bharat", "Russian", "Thailand", "Japan", "korian"} // slice
	printType(countries)

	var mp = map[string]interface{}{"Name": "Chandu", "Roll": 12, "Class": "MCA", "College": "SRM", "Fees": 100000.390} // map
	printType(mp)

	var arr = [3]int{12, 34, 67} // Array
	printType(arr)
}
