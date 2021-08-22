package main

import (
	"fmt"
	"net"
)

func main() {
	//1.连接服务器
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 9090,
	})
	if err != nil {
		fmt.Printf("connect UDP failed,err:%v\n", err)
		return
	}

	for i := 0; i < 100; i++ {
		//2.发送数据
		_, err = conn.Write([]byte("hello server!"))
		if err != nil {
			fmt.Printf("Write failed,err:%v\n", err)
			return
		}
		//3.接受数据
		var data [1024]byte
		n, remoteAddr, err := conn.ReadFromUDP(data[:])
		if err != nil {
			fmt.Printf("ReadFromUDP failed,err:%v\n", err)
			return
		}
		fmt.Printf("addr:%v,data:%v,count:%v\n", remoteAddr, string(data[:n]), n)
	}

}
