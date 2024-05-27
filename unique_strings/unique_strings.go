// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

package main

import "fmt"

func main() {
	strings := []string{"cat", "cat", "dog", "cat", "tree"}
	uniqueStrings := make(map[string]struct{})

	for _, s := range strings {
		uniqueStrings[s] = struct{}{}
	}

	for k := range uniqueStrings {
		fmt.Println(k)
	}
}
