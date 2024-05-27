// Реализовать пересечение двух неупорядоченных множеств.
package main

import "fmt"

func intersection(set1, set2 map[int]struct{}) []int {
	var result []int
	for k := range set1 {
		if _, exists := set2[k]; exists {
			result = append(result, k)
		}
	}
	return result
}

func main() {
	set1 := map[int]struct{}{1: {}, 2: {}, 3: {}, 4: {}}
	set2 := map[int]struct{}{3: {}, 4: {}, 5: {}, 6: {}}

	result := intersection(set1, set2)
	fmt.Println(result)
}
