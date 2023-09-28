package main

import "fmt"

// this is the first Structure
type rectagle struct {
	length    int
	breath    int
	geometory struct { // nested structure
		area      int
		perimeter int
	}
}

func abc() {
	rect := rectagle{
		length: 23,
		breath: 48,
	}
	rect.geometory.area = rect.length * rect.breath
	rect.geometory.perimeter = 2 * (rect.length + rect.breath)
	fmt.Println(rect)
	fmt.Println("Area =", rect.geometory.area)
	fmt.Println("Perimeter =", rect.geometory.perimeter)
}

// this is the second structure
type rectagles struct {
	length  int
	breadth int
	color   string
}

// method of the second structure
func abc1() {
	rect1 := rectagles{
		length:  23,
		breadth: 89,
		color:   "Red",
	}
	fmt.Println("Rect1 =", rect1)
	rect2 := rectagles{
		length:  48,
		breadth: 0,
		color:   "Green",
	} // Here, breadth is skipped
	fmt.Println("Rect2 = ", rect2)

	rect3 := rectagles{
		length:  0, // Here, length is skipped
		breadth: 39,
		color:   "yellow",
	}
	fmt.Println("Rect3 = ", rect3)
	rect4 := rectagles{
		length:  39,
		breadth: 35,
		color:   "", // color is skipped
	}
	fmt.Println("Rect4 = ", rect4)
}

// third structure
type str struct {
	length  int
	breadth int
	color   string
}

// third method of the third structure
func abc2() {
	st := new(str) // st is a pointer to an instance of str
	st.length = 38
	st.breadth = 90
	st.color = "Orange"
	fmt.Println("Pointer instance = ", st)
	var st1 = new(str) // st1 is an instance of str
	st1.length = 74
	st1.color = "Yellow"
	fmt.Println("instance variable = ", st1)
}
func main() {
	//abc()
	//abc1()
	abc2()

}
