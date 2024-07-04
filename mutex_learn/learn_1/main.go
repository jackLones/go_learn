package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter Counter
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 10000; j++ {
				counter.Incr()

			}
		}()
	}
	wg.Wait()
	fmt.Println(counter.count())
}

type Counter struct {
	mu    sync.Mutex
	Count int
}

func (c *Counter) Incr() {
	c.mu.Lock()
	c.Count++
	c.mu.Unlock()
}

func (c *Counter) count() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.Count
}

//func main2() {
//	var counter Counter
//	var wg sync.WaitGroup
//	wg.Add(10)
//	for i := 0; i < 10; i++ {
//		go func() {
//			defer wg.Done()
//			for j := 0; j < 10000; j++ {
//				counter.Lock()
//				counter.Count++
//				counter.Unlock()
//			}
//		}()
//	}
//	wg.Wait()
//	fmt.Println(counter.Count)
//}

//type Counter struct {
//	sync.Mutex
//	Count int
//}

func main1() {
	var mu sync.Mutex
	var count = 0
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 10000; j++ {
				mu.Lock()
				count++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println(count)
}
