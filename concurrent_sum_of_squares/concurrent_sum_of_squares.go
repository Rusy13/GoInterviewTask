// Дана последовательность чисел: 2,4,6,8,10.
// Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.

package main

import (
	"fmt"
	"sync"
)

func square(n int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- n * n
}

func main() {
	nambers := []int{2, 4, 6, 8, 10}
	ch := make(chan int, len(nambers))
	var wg sync.WaitGroup
	for _, val := range nambers {
		wg.Add(1)
		go square(val, ch, &wg)
	}
	wg.Wait()
	close(ch)

	sum := 0
	for sq := range ch {
		sum += sq
	}
	fmt.Printf("Sum of squares is %d\n", sum)
}
