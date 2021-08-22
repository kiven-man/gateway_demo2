package main

import (
	"fmt"
	"gateway_demo/demo/base/unpack/unpack"
	"net"
)

func main() {
	//1.连接服务器
	conn, err := net.Dial("tcp", "localhost:9090")
	defer conn.Close()
	if err != nil {
		fmt.Printf("connect failed,err:%v\n", err.Error())
		return
	}
	//2.解码数据
	unpack.Encode(conn, "hello world 0!!!")
}
