package main

import (
	"fmt"
	"net"
)

func main() {
	//1.监听端口
	listener, err := net.Listen("tcp", "0.0.0.0:9090")
	if err != nil {
		fmt.Printf("监听错误：%v\n", err)
		return
	}
	//2.建立套接字连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("建立套接字连接错误：%v\n", err)
			continue
		}
		//3.创建协程
		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	for {
		var buf [128]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Printf("Read错误：%v\n", err)
			break
		}
		str := string(buf[:n])
		fmt.Printf("接收数据：%v\n", str)
	}
}
