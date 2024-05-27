// Разработать программу, которая проверяет, что все символы в строке уникальные
// (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.

// Например:
// abcd — true
// abCdefAaf — false
// 	aabcd — false

package main

import (
	"fmt"
	"strings"
)

// AreAllCharactersUnique проверяет, что все символы в строке уникальны.
func AreAllCharactersUnique(s string) bool {
	// Приводим строку к нижнему регистру
	s = strings.ToLower(s)
	// Используем mapУ для отслеживания встреченных символов
	charMap := make(map[rune]bool)

	// Проверяем каждый символ
	for _, char := range s {
		// Если символ уже есть в карте, возвращаем false
		if charMap[char] {
			return false
		}
		// Иначе добавляем символ в карту
		charMap[char] = true
	}
	return true
}

func main() {
	tests := []string{
		"abcd",      // true
		"abCdefAaf", // false
		"aabcd",     // false
	}

	for _, test := range tests {
		fmt.Printf("Are all characters in '%s' unique? %t\n", test, AreAllCharactersUnique(test))
	}
}
