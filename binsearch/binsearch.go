// Реализовать бинарный поиск встроенными методами языка.

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	index := sort.SearchInts(nums, 5)
	fmt.Println("Found at index:", index)
}
