package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

func main() {
	try()
}

// 复制Mutex定义的常量
const (
	mutexLocked      = 1 << iota // 加锁标识位置
	mutexWoken                   // 唤醒标识位置
	mutexStarving                // 锁饥饿标识位置
	mutexWaiterShift = iota      // 标识waiter的起始bit位
)

// Mutex 扩展一个Mutex结构
type Mutex struct {
	sync.Mutex
}

// 尝试获取锁
func TryLock(M *Mutex) bool {
	// 如果能成功抢到锁
	if atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&M.Mutex)), 0, mutexLocked) {
		return true
	}
	// 如果处于唤醒、加锁或者饥饿状态，这次请求就不参与竞争了，返回false
	old := atomic.LoadInt32((*int32)(unsafe.Pointer(&M.Mutex)))
	if old&(mutexWoken|mutexLocked|mutexStarving) != 0 {
		return false
	}

	// 尝试在竞争的状态下请求锁
	new := old | mutexLocked
	return atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&M.Mutex)), old, new)
}
func try() {
	//程序运行时会启动一个 goroutine 持有锁，经过随机的时间才释放。主 goroutine 会尝试获取这把锁。如果前一个 goroutine 一秒内释放了这把锁，那么，主 goroutine 就有可能获取到这把锁了，输出“got the lock”，否则没有获取到也不会被阻塞，会直接输出“
	var mu Mutex
	go func() {
		mu.Lock()
		defer mu.Unlock()
		time.Sleep(time.Duration(rand.Intn(2)) * time.Second)
	}()
	time.Sleep(time.Second)
	//开始尝试获取锁
	ok := mu.TryLock()
	if ok {
		fmt.Println("got the lock")
		// do something
		mu.Unlock()
		return
	}
	// 没有获取到
	fmt.Println("can't get the lock")
}
