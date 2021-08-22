package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

func main() {
	//1.创建连接池
	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConns:          100,              //最大空闲连接
		IdleConnTimeout:       90 * time.Second, //空闲超时时间
		TLSHandshakeTimeout:   10 * time.Second, //TLS握手超时时间
		ExpectContinueTimeout: 1 * time.Second,  //100-continue状态码超时时间
	}
	//2.创建客户端
	client := http.Client{
		Timeout:   30 * time.Second, //请求超时时间
		Transport: transport,
	}
	//3.请求数据
	resp, err := client.Get("http://127.0.0.1:1210/bye")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	//4.读取内容
	bt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bt))
}
