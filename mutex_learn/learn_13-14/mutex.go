package main

import (
	"fmt"
	"time"
)

// 使用chan实现互斥锁
type Mutex struct {
	ch chan struct{}
}

// 使用锁需要初始化
func NewMutex() *Mutex {
	mu := &Mutex{
		ch: make(chan struct{}, 1),
	}
	mu.ch <- struct{}{}
	return mu
}

// 请求锁，直到获取到
func (m *Mutex) Lock() {
	<-m.ch
}

// 解锁
func (m Mutex) Unlock() {
	select {
	case m.ch <- struct{}{}:
	default:
		panic("unlock unlocked mutex")
	}
}

// 尝试获取锁
func (m *Mutex) TryLock() bool {
	select {
	case <-m.ch:
		return true
	default:
	}
	return false
}

// 加入一个超时的设置
func (m Mutex) LockTimeout(timeout time.Duration) bool {
	timer := time.NewTimer(timeout)
	select {
	case <-m.ch:
		timer.Stop()
		return true
	case <-timer.C:
	}
	return false
}
func main() {
	m := NewMutex()
	ok := m.TryLock()
	fmt.Println("locked v %v\n", ok)
	ok = m.TryLock()
	fmt.Println("locked %v\n", ok)
}
