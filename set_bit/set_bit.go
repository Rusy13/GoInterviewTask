// Дана переменная int64.
// Разработать программу которая устанавливает i-й бит в 1 или 0.

package main

import "fmt"

// SetBit устанавливает i-й бит в 1 или 0
func SetBit(n int64, pos uint, value int) int64 {
	if value == 0 {
		n &^= (1 << pos)
	} else {
		n |= (1 << pos)
	}
	return n
}

func main() {
	var n int64 = 5 // 0101 in binary
	var pos uint = 1
	fmt.Printf("Original: %b\n", n)

	n = SetBit(n, pos, 1)
	fmt.Printf("Set bit %d to 1: %b\n", pos, n)

	n = SetBit(n, pos, 0)
	fmt.Printf("Set bit %d to 0: %b\n", pos, n)
}
