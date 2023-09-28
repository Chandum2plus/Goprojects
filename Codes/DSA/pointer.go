package main

import "fmt"

type SS struct {
	a int
}

func (s *SS) abc() {
	s.a = 67
}
func main() {
	var st SS
	st.a = 23
	st.abc()
	fmt.Print(st.a)
}
