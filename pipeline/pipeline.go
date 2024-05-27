// Разработать конвейер чисел.
// Даны два канала: в первый пишутся числа (x) из массива,
// во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.

package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4, 5}
	ch1 := make(chan int)
	ch2 := make(chan int)

	// Пишем числа в первый канал
	go func() {
		for _, n := range nums {
			ch1 <- n
		}
		close(ch1)
	}()

	// Читаем из первого канала и пишем во второй
	go func() {
		for n := range ch1 {
			ch2 <- n * 2
		}
		close(ch2)
	}()

	// Читаем из второго канала и выводим в stdout
	for res := range ch2 {
		fmt.Println(res)
	}
}
