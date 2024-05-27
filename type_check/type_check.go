// Разработать программу, которая в рантайме способна определить тип переменной:
// int, string, bool, channel из переменной типа interface{}.

package main

import "fmt"

func checkType(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("Type is int")
	case string:
		fmt.Println("Type is string")
	case bool:
		fmt.Println("Type is bool")
	case chan int:
		fmt.Println("Type is channel")
	default:
		fmt.Println("Unknown type")
	}
}

func main() {
	checkType(42)
	checkType("hello")
	checkType(true)
	ch := make(chan int)
	checkType(ch)
}
