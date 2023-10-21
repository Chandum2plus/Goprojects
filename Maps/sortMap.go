package main

import (
	"fmt"
	"sort"
)

// sorting the map using the key
func main() {
	m := map[int]string{2: "Chandu", 1: "Paras", 5: "Saurabh", 4: "Sonal"}
	key := make([]int, 0, len(m))

	for k := range m {
		key = append(key, k)
	}

	sort.Ints(key)
	for _, r := range key {
		fmt.Println(r, m[r])
	}
}
