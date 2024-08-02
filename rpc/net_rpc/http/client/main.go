package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	// 1. 用 rpc 链接服务器 --Dial()
	conn, err := rpc.DialHTTP("tcp", "127.0.0.1:8800")
	if err != nil {
		fmt.Println("Dial err:", err)
		return
	}

	defer conn.Close()

	// 2. 调用远程函数
	var reply1 string // 接受返回值 --- 传出参数
	err1 := conn.Call("World.HelloWorld", "张三", &reply1)
	if err1 != nil {
		fmt.Println("Call:", err1)
		return
	}

	fmt.Println(reply1)

	var reply2 string // 接受返回值 --- 传出参数
	err2 := conn.Call("World.Print", "李四", &reply2)
	if err2 != nil {
		fmt.Println("Call:", err2)
		return
	}
	fmt.Println(reply2)
}
