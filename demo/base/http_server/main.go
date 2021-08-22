package main

import (
	"log"
	"net/http"
	"time"
)

var (
	addr = "0.0.0.0:1210"
)

func main() {
	//1.创建路由器
	mux := http.NewServeMux()
	//2.设置路由规则
	mux.HandleFunc("/bye", sayBye)
	//3.创建服务器
	server := &http.Server{
		Addr:         addr,
		WriteTimeout: time.Second * 3,
		Handler:      mux,
	}
	//4.监听端口并进行服务
	log.Println("starting httpServer at" + addr)
	err := server.ListenAndServe()
	log.Fatal(err)
}

func sayBye(w http.ResponseWriter, r *http.Request) {
	time.Sleep(1 * time.Second)
	w.Write([]byte("bye bye ,this is httpServer"))
}
