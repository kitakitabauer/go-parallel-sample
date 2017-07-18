package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func waitGroup() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			n := rand.Intn(5)
			fmt.Println("Going to sleep for", n, "seconds")
			time.Sleep(time.Duration(n) * time.Second)
		}()
	}

	wg.Wait()
	fmt.Println("All goroutines finished")
}

func main() {
	waitGroup()
}
