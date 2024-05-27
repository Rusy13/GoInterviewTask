// Разработать программу, которая переворачивает подаваемую на ход строку
// (например: «главрыба — абырвалг»). Символы могут быть unicode.

package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	words := strings.Split(s, " ")
	var result []string
	for i := len(words) - 1; i >= 0; i-- {
		result = append(result, words[i])
	}
	return strings.Join(result, " ")
}

func main() {
	s := "snow dog sun"
	fmt.Println(reverseWords(s))
}
