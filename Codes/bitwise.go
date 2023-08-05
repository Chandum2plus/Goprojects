package main

import (
	"fmt"
)

func bit() {
	var a uint = 60 /* 60 = 0011 1100 */
	var b uint = 13 /* 13 = 0000 1101 */
	var c uint = 0

	c = a & b /* 12 = 0000 1100 */
	fmt.Printf("value AND Operator is = %d\n", c)

	c = a | b /* 61 = 0011 1101 */
	fmt.Printf("value OR Operator is = %d\n", c)

	c = a ^ b /* 49 = 0011 0001 */
	fmt.Printf("Value XOR Operator is = %d\n", c)

	c = a << 2 /* 240 = 1111 0000 */
	fmt.Printf("Value LEFT SHIFT is = %d\n", c)

	c = a >> 3 /* 15 = 0000 1111 */
	fmt.Printf("Value RIGHT SHIFT is = %d\n", c)
}
func main() {
	bit()
}
