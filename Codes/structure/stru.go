package main

import "fmt"

type Rectangle struct {
	length float64
	breath float64
	color  string
}
type Vertex struct {
	x int
	y int
}

func structure() {
	fmt.Println(Rectangle{23.45, 67.09, "Red"})
	// struct fields are accessed using by a dot
	r := Rectangle{
		length: 23.78,
		breath: 45.57,
		color:  "Orange",
	}
	Res := r.length * r.breath
	fmt.Println(Res)
	fmt.Println(r.color)
	v := Vertex{}
	v.y = 48
	v.x = 47
	fmt.Println("value of x ", v.x)
	fmt.Println("value of y ", v.y)
	//Struct fields can be accessed through a struct pointer.
	//Struct fields can be accessed through a struct pointer.
	//To access the field X of a struct when we have the struct pointer p we could write (*p).X.
	//However, that notation is cumbersome, so the language permits us instead to write just p.X, without the explicit dereference.
	p := &v     // p is stored the address of the v, which v is instance variable of Vertex structure
	p.y = 90    // here, 90 is assigned in y that's why y has 90
	(*p).x = 89 //
	//	p.x = 56
	fmt.Println("value of y pointer ", v)
	fmt.Println("value of x pointer ", v)
	var (
		v1 = Vertex{1, 3}  // type has vertex
		v2 = Vertex{x: 1}  // y:0 is implicit
		v3 = Vertex{}      // x:0 and y:0
		v4 = &Vertex{1, 4} // has type *vertex
	)
	fmt.Println(v1, v2, v3, v4)
}

func main() {
	structure()
}
