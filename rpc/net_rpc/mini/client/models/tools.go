package models

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

var serverName = "HelloService"

type RPCClient struct {
	Client *rpc.Client
	Conn   net.Conn
}

func NewRpcClient(addr string) RPCClient {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println("链接服务器失败")
		return RPCClient{}
	}
	//套接字和rpc服务绑定
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	return RPCClient{Client: client, Conn: conn}
}

func (this *RPCClient) CallFunc(req string, resp *string) error {
	return this.Client.Call(serverName+".HelloWorld", req, resp)
}
