package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
)

type proxy struct{}

func (p *proxy) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	fmt.Printf("Received request %s %s %s\n", req.Method, req.Host, req.RemoteAddr)
	//定义默认tranceport
	tranceport := http.DefaultTransport
	//1.浅拷贝Request
	outReq := new(http.Request)
	*outReq = *req
	//2.获取clientIP，并追加到X-Forward-For
	if clientIP, _, err := net.SplitHostPort(req.RemoteAddr); err == nil {
		if prior, ok := outReq.Header["X-Forwarded-For"]; ok {
			clientIP = strings.Join(prior, ",") + "," + clientIP
		}
		outReq.Header.Set("X-Forwarded-For", clientIP)
	}
	//3.向下游请求
	res, err := tranceport.RoundTrip(outReq)
	if err != nil {
		rw.WriteHeader(http.StatusBadGateway)
	}
	//4.拷贝head
	for key, val := range res.Header {
		for _, v := range val {
			rw.Header().Add(key, v)
		}
	}
	//5.拷贝Body
	io.Copy(rw, res.Body)
	res.Body.Close()
}
func main() {
	fmt.Println("Serve on :8080")
	//设置路由
	http.Handle("/", &proxy{})
	//设置监听并开启服务
	http.ListenAndServe("0.0.0.0:8080", nil)
}
