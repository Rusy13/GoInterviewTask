// Удалить i-ый элемент из слайса.

package main

import "fmt"

func removeElement(s []int, i int) []int {
	return append(s[:i], s[i+1:]...)
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	indexToRemove := 2
	nums = removeElement(nums, indexToRemove)
	fmt.Println(nums)
}
