package main

import (
	"context"
	"fmt"
	"strings"
	"time"
)

func Job(ctx context.Context, input chan string) chan string {
	output := make(chan string)
	go func() {
		defer close(output)
		for {
			select {
			case <-ctx.Done():
				fmt.Println("cancel:", ctx.Err())
				return
			case v, ok := <-input:
				if !ok {
					fmt.Println("came input")
					return
				}
				output <- strings.ToUpper(v)
			}
		}
	}()
	return output
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	defer cancel()
	input := make(chan string)
	go func() {
		defer close(input)
		for _, v := range []string{"hoge", "moge"} {
			select {
			case <-ctx.Done():
				return
			case input <- v:
				//
			}
		}
		time.Sleep(1100 * time.Millisecond)
	}()
	for v := range Job(ctx, input) {
		fmt.Println(v)
	}
	time.Sleep(100 * time.Millisecond)
}
