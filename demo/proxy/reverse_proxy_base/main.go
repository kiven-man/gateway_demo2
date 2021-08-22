package main

import (
	"bufio"
	"log"
	"net/http"
	"net/url"
)

var (
	proxy_addr = "http://127.0.0.1:2003"
	port       = "2002"
)

func handler(rw http.ResponseWriter, req *http.Request) {
	//1.解析代理地址，并更改请求体的协议和主机
	proxy, _ := url.Parse(proxy_addr)
	req.URL.Scheme = proxy.Scheme
	req.URL.Host = proxy.Host

	//2.请求下游
	transpoat := http.DefaultTransport
	res, err := transpoat.RoundTrip(req)
	if err != nil {
		log.Fatal(err)
		return
	}
	//3.拷贝header
	for key, val := range res.Header {
		for _, v := range val {
			rw.Header().Add(key, v)
		}
	}
	//4.拷贝body
	defer res.Body.Close()
	bufio.NewReader(res.Body).WriteTo(rw)
	//io.Copy(rw, res.Body)

}

func main() {
	log.Println("Start serving on port " + port)
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
