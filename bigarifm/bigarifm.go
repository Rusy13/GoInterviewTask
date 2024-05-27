// Разработать программу, которая перемножает, делит,
// складывает, вычитает две числовых переменных a,b, значение которых > 2^20.

package main

import "fmt"

func main() {
	a, b := 2<<20, 3<<20
	fmt.Printf("a + b = %d\n", a+b)
	fmt.Printf("a - b = %d\n", a-b)
	fmt.Printf("a * b = %d\n", a*b)
	fmt.Printf("a / b = %d\n", a/b)
}
