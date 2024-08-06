package main

//grpc客户端代码

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc/client/Hello/proto/helloService"
)

func main() {
	// 1、连接服务器
	/*
	   credentials.NewClientTLSFromFile ：从输入的证书文件中为客户端构造TLS凭证。
	   grpc.WithTransportCredentials ：配置连接级别的安全凭证（例如，TLS/SSL），返回一个
	   DialOption，用于连接服务器。
	*/
	grpcClient, err := grpc.Dial("127.0.0.1:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)
	}

	//2、注册客户端
	client := helloService.NewHelloClient(grpcClient)
	//3、调用服务端函数, 实现HelloClient接口:SayHello()
	/*
	   // HelloClient is the client API for Hello service.
	   //
	   // For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
	   type HelloClient interface {
	       // 通过rpc来指定远程调用的方法:
	       // SayHello方法, 这个方法里面实现对传入的参数HelloReq, 以及返回的参数HelloRes进行约束
	       SayHello(ctx context.Context, in *HelloReq, opts ...grpc.CallOption) (*HelloRes, error)
	   }
	*/
	res, err1 := client.SayHello(context.Background(), &helloService.HelloReq{
		Name: "张三",
	})
	if err1 != nil {
		fmt.Printf("调用服务端代码失败: %s", err1)
		return
	}

	fmt.Printf("%#v\r\n", res)
	fmt.Printf("调用成功: %s", res.Message)
}
