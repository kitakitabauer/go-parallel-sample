package main

import (
	"fmt"
	"strconv"
)

func getInt() int {
	return 1
}

func getMap() map[string]string {
	return map[string]string{
		"_id": "ID1",
	}
}

func main() {
	fmt.Println("main function start")

	for i := 0; i < 3; i++ {
		intCh := make(chan int)
		mapCh := make(chan map[string]string)

		go func(i int) {
			intCh <- i
		}(i)
		go func(i int) {
			mapCh <- map[string]string{
				"i": strconv.Itoa(i),
			}
		}(i)

		i := <-intCh
		fmt.Printf("loop: %d\n", i)
		m := <-mapCh
		fmt.Printf("map: %#v\n", m)
	}

	fmt.Println("main function end")
}
