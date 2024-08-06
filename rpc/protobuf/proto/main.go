package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"proto/protoService"
)

func main() {
	//初始化并赋值
	u := &protoService.Userinfo{
		Name:  "zhangsan",
		Age:   20,
		Hobby: []string{"吃饭", "睡觉", "写代码"},
	}
	fmt.Println(u.GetHobby())
	// proto.Marshald对protoBufer进行序列化
	data, err1 := proto.Marshal(u)
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println(data)

	//proto.Unmarshal可以对protoBufer进行反序列化
	info := protoService.Userinfo{}
	err2 := proto.Unmarshal(data, &info)
	if err2 != nil {
		fmt.Println(err2)
	}

	fmt.Printf("%#v", info)
	fmt.Println(info.GetHobby())
}
