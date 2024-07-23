package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var counter Counter
	for i := 0; i < 10; i++ {
		go func() {
			for {
				fmt.Printf("count = %d", counter.Count())
				time.Sleep(time.Millisecond) //sleep 1 毫秒
			}
		}()
	}

	for {
		counter.Incr()
		time.Sleep(time.Second)
	}
}

// 一个线程安全的计数器
type Counter struct {
	mu    sync.RWMutex
	count int
}

func (c *Counter) Incr() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

// 使用读锁保护
func (c Counter) Count() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.count
}
