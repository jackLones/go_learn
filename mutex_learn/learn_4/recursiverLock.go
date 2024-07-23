package main

import (
	"fmt"
	"github.com/petermattis/goid"
	"sync"
	"sync/atomic"
)

type RecursivMutex struct {
	sync.Mutex
	owner     int64 // 当前持有锁的goroutine id
	recursion int32 // 这个goroutine 重入的次数
}

func (m *RecursivMutex) Lock() {
	gid := goid.Get()
	// 如果当前持有锁的goroutine就是这次调用的goroutine，说明是重入
	if atomic.LoadInt64(&m.owner) == gid {
		m.recursion++
		return
	}

	//m.Lock()
	m.Mutex.Lock()
	// 获得锁的goroutine是第一次调用，记录下goroutine id，调用次数加1
	atomic.StoreInt64(&m.owner, gid)
	m.recursion = 1
}

func (m *RecursivMutex) Unlock() {
	gid := goid.Get()
	// 非持有锁的goroutine尝试释放锁，抛出异常
	if atomic.LoadInt64(&m.owner) != gid {
		panic(fmt.Sprintf("wrong the owner(%d): %d!", m.owner, gid))
	}

	// 调用次数减1
	m.recursion--
	if m.recursion != 0 { // 锁没有完全释放，直接返回
		return
	}

	// 此goroutine已经完全释放锁，重置owner和recursion
	atomic.StoreInt64(&m.owner, -1)
	m.Mutex.Unlock()
}
