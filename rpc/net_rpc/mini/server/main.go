package main

import (
	"fmt"
	"mini/server/models"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// 定义类对象
type World struct{}

// 绑定类方法
func (this *World) HelloWorld(req string, res *string) error {
	fmt.Println(req)
	*res = req + " 你好!"
	return nil
	//return errors.New("未知的错误!")
}

func main() {
	//注册rpc服务 维护一个hash表，key值是服务名称，value值是服务的地址
	// rpc.RegisterName("HelloService", new(World))
	models.RegisterService(new(World))
	//设置监听
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	for {
		//接收连接
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		//给当前连接提供针对json格式的rpc服务
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
