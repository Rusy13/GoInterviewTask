// Написать программу, которая конкурентно рассчитает значение квадратов
// чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

package main

import (
	"fmt"
	"sync"
)

func squares(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Square of %d is %d\n", n, n*n)
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup

	for _, i := range numbers {
		wg.Add(1)
		go squares(i, &wg)
	}

	wg.Wait()
}
