package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Worker функция воркера
func Worker(id int, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for n := range ch {
		fmt.Printf("Worker %d received %d\n", id, n)
	}
}

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup

	// Определение количества воркеров
	numWorkers := 3
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go Worker(i, ch, &wg)
	}

	// Генерация данных
	go func() {
		for i := 0; ; i++ {
			ch <- i
			time.Sleep(500 * time.Millisecond)
		}
	}()

	// Обработка сигналов завершения
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	<-c

	close(ch)
	wg.Wait()
	fmt.Println("All workers finished")
}
