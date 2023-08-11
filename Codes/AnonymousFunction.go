package main

import "fmt"

func Name() func(p, q string) string {
	//fmt.Println(i("Chandu ", "Kumar "))
	myf := func(i, j string) string {
		return i + j + " M2-Plus"
	}
	return myf
}
func main() {
	//value := func(p, q string) string {
	//return p + q + "Pandit"
	value := Name()
	fmt.Println(value("welcome", " to"))

}
