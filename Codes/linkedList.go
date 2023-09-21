package main

import "fmt"

// example of the linked list
type Node struct {
	value int
	next  *Node
}

func linkedList() {
	head := &Node{value: 10}
	head.next = &Node{value: 20}
	head.next.next = &Node{value: 30}

	// Accessing the element from the linked list
	fmt.Println("Accessing the element of the linked list")

	node := head
	for node != nil {
		fmt.Println(node.value)
		node = node.next
	}
}
func test() {
	num := 9
	num2 := 9
	Sum := num + num2
	fmt.Println("sum of the two number is = ", Sum)
	var a, b = 7, "hello Bharat"
	fmt.Println(a)
	fmt.Println(b)

	//how to use the constant type
	const (
		Name = string("Chandu")
		N    = int(12)
		N2   = int(23)
		summ = N + N2
	)

	fmt.Println("string constant", Name)
	fmt.Println("integer constant - ", N)
	fmt.Println("integer constant - ", N2)
	fmt.Println("sum of the both constant -", summ)

}
func main() {
	//linkedList()
	test()

}
