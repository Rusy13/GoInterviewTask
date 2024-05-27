package main

import (
	"fmt"
	"sync"
)

type SafeMap struct {
	mu sync.Mutex
	m  map[int]int
}

func NewSafeMap() *SafeMap {
	return &SafeMap{m: make(map[int]int)}
}

func (sm *SafeMap) Set(key, value int) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.m[key] = value
}

func main() {
	sm := NewSafeMap()

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			sm.Set(i, i*2)
		}(i)
	}

	wg.Wait()

	sm.mu.Lock()
	defer sm.mu.Unlock()
	for key, value := range sm.m {
		fmt.Printf("key: %d, value: %d\n", key, value)
	}
}
