package main

import (
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"time"
)

func main() {
	var g errgroup.Group

	//第一个子任务，执行成功
	g.Go(func() error {
		time.Sleep(1 * time.Second)
		fmt.Println("第一个子任务执行成功")
		return nil
	})

	//第二个子任务，执行失败
	g.Go(func() error {
		time.Sleep(10 * time.Second)
		fmt.Println("第二个子任务")
		return errors.New("第二个子任务执行失败了")
	})

	//第三个子任务，执行成功
	g.Go(func() error {
		time.Sleep(15 * time.Second)
		fmt.Println("第三个子任务执行成功")
		return nil
	})

	//等待三个任务都完成
	if err := g.Wait(); err != nil {
		fmt.Println("任务执行失败：", err)
	} else {
		fmt.Println("所有任务执行成功")
	}
}
