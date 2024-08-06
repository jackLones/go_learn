package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc/server/hello/proto/helloService"
	"log"
	"net"
)

//grpc远程调用的接口,需要实现hello.proto中定义的Hello服务接口,以及里面的方法
//1.定义远程调用的结构体和方法,这个结构体需要实现HelloServer的接口

type Hello struct {
	helloService.UnimplementedHelloServer
}

// SayHello SayHello方法参考hello.pb.go中的接口
/*
// HelloServer is the server API for Hello service.
type HelloServer interface {
    // 通过rpc来指定远程调用的方法:
    // SayHello方法, 这个方法里面实现对传入的参数HelloReq, 以及返回的参数HelloRes进行约束
    SayHello(context.Context, *HelloReq) (*HelloRes, error)
}
*/
func (this Hello) SayHello(c context.Context, req *helloService.HelloReq) (*helloService.HelloRes, error) {
	fmt.Println(req)
	return &helloService.HelloRes{
		Message: "你好" + req.Name,
	}, nil
}

func main() {
	// 初始一个 grpc 对象
	grpcServer := grpc.NewServer()

	// 注册服务
	helloService.RegisterHelloServer(grpcServer, &Hello{})

	// 设置监听，指定 IP、port
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 启动服务
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
