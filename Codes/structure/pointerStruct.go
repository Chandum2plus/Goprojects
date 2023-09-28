package main

import "fmt"

type rectangle struct {
	lenght  int
	breadth int
	color   string
}

func avc() {
	var rect = &rectangle{12, 45, "Red"} // can't skipped value hare
	fmt.Println("Rect = ", rect)

	var rect2 = &rectangle{} //Here, we can skip value
	rect2.lenght = 39
	//rect2.lenght = 90
	rect2.color = "Green"
	fmt.Println("Rect2 =", rect2) // length's value is skipped

	var rect3 = &rectangle{}
	(*rect3).lenght = 48
	(*rect3).color = "Yellow"
	fmt.Println("Rect3 -", rect3) // Here, breadth is skipped
}
func main() {
	avc()

}
