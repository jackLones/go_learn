package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
)

// 定义类对象
type World struct {
}

// 绑定类方法
func (this *World) HelloWorld(req string, res *string) error {
	*res = req + " 你好!"
	return nil
	//return errors.New("未知的错误!")
}

// 绑定类方法
func (this *World) Print(req string, res *string) error {
	*res = req + " this is Print!"
	return nil
	//return errors.New("未知的错误!")
}

func main() {
	// 1. 注册RPC服务
	rpc.Register(new(World)) // 注册rpc服务
	rpc.HandleHTTP()         // 采用http协议作为rpc载体
	// 2. 设置监听
	lis, err := net.Listen("tcp", "127.0.0.1:8800")
	if err != nil {
		log.Fatalln("fatal error: ", err)
	}
	fmt.Fprintf(os.Stdout, "%s", "start connection")
	// 3. 建立链接
	http.Serve(lis, nil)
}
