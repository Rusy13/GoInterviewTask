// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{5, 3, 8, 6, 2, 7, 1, 4}
	sort.Ints(nums)
	fmt.Println(nums)
}
