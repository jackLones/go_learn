package models

import "net/rpc"

var serverName = "HelloService"

type RPCInterface interface {
	HelloWorld(string, *string) error
}

// 调用该方法时, 需要给 i 传参, 参数应该是 实现了 HelloWorld 方法的类对象!
func RegisterService(i RPCInterface) {
	rpc.RegisterName(serverName, i)
}
