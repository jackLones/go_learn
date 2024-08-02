package main

import (
	"fmt"
	"reflect"
	"time"
)

func or2(channels ...<-chan interface{}) <-chan interface{} {
	//特殊情况，只有0个或1个chan
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}

	orDOne := make(chan interface{})
	go func() {
		defer close(orDOne)
		//利用反射构建SelectCase
		var cases []reflect.SelectCase
		for _, v := range channels {
			cases = append(cases, reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(v),
			})
		}
		//随机选一个可以的case
		reflect.Select(cases)
	}()
	return orDOne
}
func or1(channels ...<-chan interface{}) <-chan interface{} {
	//特殊情况，只有0个或1个chan
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}

	orDOne := make(chan interface{})
	for _, ch := range channels {
		go func(ch <-chan interface{}) {
			select {
			case <-ch:
				orDOne <- struct{}{}
			}
		}(ch)
	}
	return orDOne
}
func or(channels ...<-chan interface{}) <-chan interface{} {
	//特殊情况，只有0个或1个chan
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}

	orDOne := make(chan interface{})
	go func() {
		defer close(orDOne)
		switch len(channels) {
		case 2: //也是一种特殊情况
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		default: // 超过两个，使用二分递归法来处理
			m := len(channels) / 2
			select {
			case <-or(channels[:m]...):
			case <-or(channels[m:]...):
			}
		}
	}()
	return orDOne
}

func main() {
	start := time.Now()
	<-or2(
		sig(10*time.Second),
		sig(20*time.Second),
		sig(30*time.Second),
		sig(40*time.Second),
		sig(50*time.Second),
		sig(1*time.Second),
	)
	fmt.Println("done after %v", time.Since(start))
}

func sig(after time.Duration) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
	}()
	return c
}
