// concept of the linked list

package main

import "fmt"

// Declaring the Node structure
type Node struct {
	value int   // this is the int type of value
	next  *Node // this is the next pointer which type is Node and it contain the address of the next node
}

// this is the add function and it has some argument Head is head of the node
// that and its type is Node, value is a integer type of variable with the
// help of  we will pass the node
func addNode(head *Node, value int) *Node {
	newNode := &Node{value: value} // Here, create the new node
	if head == nil {               // Here, checking the condition
		head = newNode
	} else {
		current := head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
	return head
}

func main() {
	fmt.Println("The element added in the linked list are -: ")
	var head *Node
	head = addNode(head, 23) // Add elements to the list
	head = addNode(head, 25)
	head = addNode(head, 26)
	head = addNode(head, 63)
	head = addNode(head, 34)
	current := head
	for current != nil {
		fmt.Println(current.value) // Print the Added value
		current = current.next
	}
}
