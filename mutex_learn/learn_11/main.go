package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second) // 修改此处超时时间，可打印出不同结果
	defer cancel()                                                          // 避免其他地方忘记cancel，且重复调用不影响

	ids := fetchWebData(ctx)

	fmt.Println(ids)
}

func fetchWebData(ctx context.Context) (res []int64) {
	select {
	case <-time.After(3 * time.Second):
		return []int64{100, 200, 300}
	case <-ctx.Done():
		return []int64{1, 2, 3}
	}
}

func main1() {
	ctx := context.Background()
	process(ctx)
	ctx = context.WithValue(ctx, "traceId", "1111")
	process(ctx)
}

func process(ctx context.Context) {
	traceId, ok := ctx.Value("traceId").(string)
	if ok {
		fmt.Printf("process over. trace_id=%s\n", traceId)
	} else {
		fmt.Printf("process over. no trace_id\n")
	}
}
