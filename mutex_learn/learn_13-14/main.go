package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 11)
	for i := 0; i < 10; i++ {
		select {
		case ch <- i:
			//case v := <-ch:
			//	fmt.Println(v)
		}
	}

	for v := range ch {
		fmt.Println(v)
	}
}
