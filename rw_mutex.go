package main

import (
	"fmt"
	"sync"
	"time"
)

func rwMutex() {
	var mu sync.RWMutex
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		id := i
		go func() {
			defer wg.Done()
			for i := 0; i < 5; i++ {
				mu.RLock()
				fmt.Printf("Reader %d: Acquired lock\n", id)
				time.Sleep(time.Second)
				mu.RUnlock()
			}
		}()
	}

	time.AfterFunc(3*time.Second, func() {
		fmt.Println("Writer: Acquired lock")
		mu.Lock()
	})
	time.AfterFunc(6*time.Second, func() {
		fmt.Println("Writer: Releasing lock")
		mu.Unlock()
	})

	wg.Wait()
}

func main() {
	rwMutex()
}
