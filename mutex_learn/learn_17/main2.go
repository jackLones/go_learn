package main

import (
	"golang.org/x/sync/singleflight"
	"log"
	"sync"
	"time"
)

var (
	sfKey1 = "key1"
	wg     *sync.WaitGroup
	sf     singleflight.Group
	nums   = 10
)

func main() {
	getValueService("key")
}

func getValueService(key string) {
	var val string
	wg = &sync.WaitGroup{}
	wg.Add(nums)
	// 模拟多协程同时请求
	for idx := 0; idx < nums; idx++ {
		go func(idx int) {
			defer wg.Done()
			value, _ := getValueBySingleflight(idx, key) //简化代码，不处理error
			log.Printf("request %v get value: %v", idx, value)
			val = value
		}(idx)
	}
	wg.Wait()
	log.Println("val: ", val)
	return
}

// getValueBySingleflight 使用singleflight取cacheKey对应的value值
func getValueBySingleflight(idx int, cacheKey string) (string, error) {
	log.Printf("idx %v into-cache...", idx)
	// 调用singleflight的Do()方法
	value, _, _ := sf.Do(cacheKey, func() (ret interface{}, err error) {
		log.Printf("idx %v is-setting-cache", idx)
		// 休眠0.1s以捕获并发的相同请求
		time.Sleep(100 * time.Millisecond)
		log.Printf("idx %v set-cache-success!", idx)
		return "myValue", nil
	})
	return value.(string), nil
}

/**
2024/08/01 14:22:35 idx 0 into-cache...
2024/08/01 14:22:35 idx 0 is-setting-cache
2024/08/01 14:22:35 idx 4 into-cache...
2024/08/01 14:22:35 idx 5 into-cache...
2024/08/01 14:22:35 idx 6 into-cache...
2024/08/01 14:22:35 idx 7 into-cache...
2024/08/01 14:22:35 idx 3 into-cache...
2024/08/01 14:22:35 idx 8 into-cache...
2024/08/01 14:22:35 idx 9 into-cache...
2024/08/01 14:22:35 idx 2 into-cache...
2024/08/01 14:22:35 idx 1 into-cache...
2024/08/01 14:22:35 idx 0 set-cache-success!
2024/08/01 14:22:35 request 0 get value: myValue
2024/08/01 14:22:35 request 9 get value: myValue
2024/08/01 14:22:35 request 8 get value: myValue
2024/08/01 14:22:35 request 5 get value: myValue
2024/08/01 14:22:35 request 4 get value: myValue
2024/08/01 14:22:35 request 6 get value: myValue
2024/08/01 14:22:35 request 7 get value: myValue
2024/08/01 14:22:35 request 2 get value: myValue
2024/08/01 14:22:35 request 3 get value: myValue
2024/08/01 14:22:35 request 1 get value: myValue
2024/08/01 14:22:35 val:  myValue

由结果可以看到，索引=1的协程第一个进入了Do()方法，其他协程则阻塞住,等到idx=1的协程拿到执行结果后，协程以乱序的形式返回执行结果。
相同key的情况下，singleflight将我们的多个请求合并成1个请求。由1个请求去执行对共享资源的操作。
*/
