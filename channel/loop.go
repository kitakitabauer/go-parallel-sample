package main

import (
	"fmt"
	"time"
)

var (
	allWorkerNum = 3
	doneNum      int

	workerCh = make(chan int, allWorkerNum)

	jobComplete = false
)

func main() {

	for i := 0; i < allWorkerNum; i++ {
		go worker()
	}

	ch := make(chan int)

	go func() {
		defer close(ch)
		for i := 1; i <= 30; i++ {
			ch <- i
			fmt.Printf(" Write - val: %d, len:%d\n", i, len(ch))
		}
	}()

	go func() {
		for v := range ch {
			fmt.Printf(" Read - val: %d, len:%d\n", v, len(ch))
			workerCh <- v
		}

		time.Sleep(5 * time.Millisecond)
		jobComplete = true
	}()

	Wait()
}

func worker() {
	for {
		select {
		case i := <-workerCh:
			fmt.Println("worker側にきた", i)
			time.Sleep(5 * time.Millisecond)
			doneNum++
		}
	}
}

func Wait() {
	for {
		if jobComplete {
			fmt.Println("finish!")
			return
		}
	}

}
