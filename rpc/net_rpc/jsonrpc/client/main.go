package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

//rpc服务端

/*
把默认的rpc 改为jsonrpc
    1、rpc.Dial需要调换成net.Dial
    2、增加建立基于json编解码的rpc服务  client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
    3、conn.Call 需要改为client.Call
*/

func main() {
	//1、用 net.Dial和rpc微服务端建立连接
	conn, err1 := net.Dial("tcp", "127.0.0.1:8080")
	if err1 != nil {
		fmt.Println(err1)
	}
	//2、当客户端退出的时候关闭连接
	defer conn.Close()

	//建立基于json编解码的rpc服务
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	//3、调用远程函数
	//微服务端返回的数据
	var reply string
	/*
	   1、第一个参数: hello.SayHello,hello 表示服务名称  SayHello 方法名称
	   2、第二个参数: 给服务端的req传递数据
	   3、第三个参数: 需要传入地址,获取微服务端返回的数据
	*/
	err2 := client.Call("hello.SayHello", "张三", &reply)
	if err2 != nil {
		fmt.Println(err2)
	}
	//4、获取微服务返回的数据
	fmt.Println(reply)
}
