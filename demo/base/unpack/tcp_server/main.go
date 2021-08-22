package main

import (
	"fmt"
	"gateway_demo/demo/base/unpack/unpack"
	"net"
)

func main() {
	//1.监听接口
	listener, err := net.Listen("tcp", "0.0.0.0:9090")
	if err != nil {
		fmt.Printf("listen fail ,err: %v\n", err)
		return
	}
	//2.接受请求
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("accept fail ,err:%v\n", err)
			continue
		}
		//3.创建独立协程
		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	for {
		bodyBuf, err := unpack.Decode(conn)
		if err != nil {
			fmt.Printf("read from conn failed ,err:%v\n", err)
			return
		}
		str := string(bodyBuf)
		fmt.Printf("recive from client ,data:%v\n", str)
	}
}
