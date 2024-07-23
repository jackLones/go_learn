package main

import (
	"context"
	"fmt"
	"time"
)

func worker() {
	fmt.Println("Goroutine started")

	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Goroutine timeout")
	case <-time.After(5 * time.Second):
		fmt.Println("Goroutine finished")
	}
}

func main() {
	go worker()
	time.Sleep(6 * time.Second)
	fmt.Println("Program finished")
}
func worker3(ctx context.Context) {
	fmt.Println("Goroutine started")

	select {
	case <-ctx.Done():
		fmt.Println("Goroutine cancelled")
		return
	default:
		fmt.Println("Goroutine finished")
	}
}

func main3() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	go worker3(ctx)
	time.Sleep(2 * time.Second)
	fmt.Println("Program finished")
}

func main2() {
	done := make(chan bool, 1)
	go func() {
		fmt.Println("Goroutine started")
		// do some work
		fmt.Println("Goroutine finished")
		done <- true
	}()

	<-done
	//time.Sleep(1 * time.Second)
	fmt.Println("Program finished")
}

func main1() {
	go func() {
		fmt.Println("Goroutine started")
		// do some work
		fmt.Println("Goroutine finished")
	}()
	//time.Sleep(1 * time.Second)
	fmt.Println("Program finished")
}
