package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	go func() {
		ch1 <- 1
	}()
	select {
	case <-ch1:
		fmt.Println(1)
	}
}
func main2() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	ch4 := make(chan int)
	go func() {
		for {
			fmt.Println("1")
			time.Sleep(1 * time.Second)
			ch2 <- 1 //I'm done, you turn
			<-ch1
		}
	}()

	go func() {
		for {
			<-ch2
			fmt.Println("2")
			time.Sleep(1 * time.Second)
			ch3 <- 1
		}

	}()

	go func() {
		for {
			<-ch3
			fmt.Println("3")
			time.Sleep(1 * time.Second)
			ch4 <- 1
		}

	}()

	go func() {
		for {
			<-ch4
			fmt.Println("4")
			time.Sleep(1 * time.Second)
			ch1 <- 1
		}

	}()

	select {}

}

type Token struct{}

func newWorker(id int, ch chan Token, nextCh chan Token) {
	for {
		token := <-ch         // 取得令牌
		fmt.Println((id + 1)) // id从1开始
		time.Sleep(time.Second)
		nextCh <- token
	}
}
func main1() {
	chs := []chan Token{make(chan Token), make(chan Token), make(chan Token), make(chan Token)}

	// 创建4个worker
	for i := 0; i < 4; i++ {
		go newWorker(i, chs[i], chs[(i+1)%4])
	}

	//首先把令牌交给第一个worker
	chs[0] <- struct{}{}

	select {}
}
