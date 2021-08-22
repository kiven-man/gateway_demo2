package main

import (
	"fmt"
	"net"
)

func main() {
	//1.监听服务器
	listener, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 9090,
	})
	if err != nil {
		fmt.Printf("ListenUDP failed,err:%v\n", err)
		return
	}
	//2.循环读取消息内容
	for {
		//读取
		var data [1024]byte
		n, addr, err := listener.ReadFromUDP(data[:])
		if err != nil {
			fmt.Printf("ReadFromUDP failed from addr:%v,err:%v\n", addr, err)
			break
		}
		//3.回复数据
		go func() {
			fmt.Printf("addr:%v,   data:%v,    count:%v\n", addr, string(data[:n]), n)
			_, err := listener.WriteToUDP([]byte("接收成功！"), addr)
			if err != nil {
				fmt.Printf("write failed,err:%v\n", err)
			}
		}()
	}

}
