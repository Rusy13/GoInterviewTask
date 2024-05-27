package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	var wg sync.WaitGroup

	// Остановка через канал
	ch := make(chan struct{})
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ch:
				fmt.Println("Stopped by channel")
				return
			default:
				// Doing some work
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	// Остановка через контекст
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Stopped by context")
				return
			default:
				// Doing some work
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	// Остановка по таймеру
	timer := time.NewTimer(2 * time.Second)
	wg.Add(1)
	go func() {
		defer wg.Done()
		select {
		case <-timer.C:
			fmt.Println("Stopped by timer")
		}
	}()

	// Остановка по сигналу операционной системы
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigCh
		fmt.Println("Stopped by OS signal")
		cancel()     // отменяем контекст
		close(ch)    // закрываем канал
		timer.Stop() // останавливаем таймер
	}()

	// Ожидание завершения всех горутин
	wg.Wait()
	fmt.Println("All goroutines stopped")
}
