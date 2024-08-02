package main

import (
	"fmt"
	"mini/client/models"
)

func main() {
	//建立tcp连接
	client := models.NewRpcClient("127.0.0.1:8080")
	//关闭连接
	defer client.Conn.Close()

	var reply string // 接受返回值 --- 传出参数
	err := client.CallFunc("this is client", &reply)

	if err != nil {
		fmt.Println("Call:", err)
		return
	}
	fmt.Println(reply)
}
