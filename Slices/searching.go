package main

import "fmt"

// Searching the item from the slice
func searching() {
	var size, serchItem, i int
	for {
		fmt.Print("Enter the size of the slice - ")
		fmt.Scan(&size)

		var serchSlic = make([]int, size)
		fmt.Print("Enter the element of the slice - ")
		for i = 0; i < size; i++ {
			fmt.Scan(&serchSlic[i])
		}
		fmt.Print("Enter the Search item - ")
		fmt.Scan(&serchItem)
		flag := 0
		for i = 0; i < size; i++ {
			if serchSlic[i] == serchItem {
				flag = 1
				break
			}
		}
		if flag == 1 {
			fmt.Println("Item is found - ", serchItem)
			fmt.Println("Position of Item - ", i)
		} else {
			fmt.Println(serchItem, " - Item is not found")
		}
	}
}
func main() {
	searching()
}
