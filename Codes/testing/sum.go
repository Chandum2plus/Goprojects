package testing

import (
	"fmt"
)

func Add(x, y int) (res int) {
	res = x + y
	fmt.Println("res - ", res)
	return res
}

func Sub(x, y int) (res int) {
	res = x - y
	fmt.Println("res - ", res)
	return res
}
func Name(name string) (res string) {
	res = "Chandu_kumar"
	fmt.Println("res - ", res)
	return res
}
func Summ(num, num2 int) (res int) {
	res = num + num2
	fmt.Println("res =", res)
	return res

}
