package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

// 复制Mutex定义的常量
const (
	mutexLocked      = 1 << iota // 加锁标识位置
	mutexWoken                   // 唤醒标识位置
	mutexStarving                // 锁饥饿标识位置
	mutexWaiterShift = iota      // 标识waiter的起始bit位
)

type Mutex struct {
	sync.Mutex
}

func (m *Mutex) Count() int {
	// 获取state字段的值
	v := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex))) //等待者的数量
	v = v>>mutexWaiterShift + (v & mutexLocked)               // 加上锁持有者的数量
	return int(v)
}

// 锁是否被持有
func (m *Mutex) IsLocked() bool {
	state := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	return state&mutexLocked == mutexLocked
}

// 是否有等待者被唤醒
func (m *Mutex) IsWoken() bool {
	state := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	return state&mutexWoken == mutexWoken
}

// 锁是否处于饥饿状态
func (m *Mutex) IsStarving() bool {
	state := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	return state&mutexStarving == mutexStarving
}

func main() {
	var mu Mutex
	for i := 0; i < 1000; i++ { // 启动1000个goroutine
		go func() {
			mu.Lock()
			time.Sleep(time.Second)
			mu.Unlock()
		}()
	}

	time.Sleep(time.Second)
	// 输出锁的信息
	fmt.Printf("waitings: %d, isLocked: %t, woken: %t,  starving: %t\n", mu.Count(), mu.IsLocked(), mu.IsWoken(), mu.IsStarving())
}
