package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

// 获取当前协程的id
func main2() {
	var buf [64]byte
	// buf[:]确保了将整个缓冲区以切片的形式传递给函数，以便函数可以往这个缓冲区写入数据而不影响原始数组的使用方式。这是一种安全且灵活的数据操作方式
	n := runtime.Stack(buf[:], false)
	// 得到id字符串
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, _ := strconv.Atoi(idField)
	//if err != nil {
	//	panic(fmt.Sprint("internal/mutex_learn/learn_3/main.go:GoID: %v", err))
	//}
	fmt.Println(id)
}

type Counter struct {
	sync.Mutex
	count int
}

func main1() {
	var c Counter
	c.Lock()
	defer c.Unlock()
	c.count++
	fool(c)

}

func fool(c Counter) {
	c.Lock()
	defer c.Unlock()
}
