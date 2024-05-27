// Разработать программу, которая будет последовательно отправлять значения в канал,
// а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.

package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	duration := 5 * time.Second

	// Генерация данных
	go func() {
		for i := 0; ; i++ {
			ch <- i
			time.Sleep(300 * time.Millisecond)
		}
	}()

	// Чтение данных с канала
	go func() {
		for n := range ch {
			fmt.Println(n)
		}
	}()

	// Завершение через N секунд
	time.Sleep(duration)
	close(ch)
	fmt.Println("Program finished")
}
