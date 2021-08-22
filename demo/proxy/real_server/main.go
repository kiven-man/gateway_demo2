package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type RealServer struct {
	Addr string
}

func main() {
	rs1 := &RealServer{Addr: "127.0.0.1:2003"}
	rs1.Run()

	rs2 := &RealServer{Addr: "127.0.0.1:2004"}
	rs2.Run()

	//监听关闭
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}

//启动函数
func (rs *RealServer) Run() {
	log.Println("Starting httpserver at " + rs.Addr)
	//路由
	mux := http.NewServeMux()
	mux.HandleFunc("/", rs.HelloHandler)
	mux.HandleFunc("/base/error", rs.ErrorHandler)
	mux.HandleFunc("/test_http_string/test_http_string/aaa", rs.TimeOutHandler)
	//服务
	server := http.Server{
		Addr:         rs.Addr,
		WriteTimeout: time.Second * 3,
		Handler:      mux,
	}

	go func() {
		log.Fatal(server.ListenAndServe())
	}()

}

func (rs *RealServer) HelloHandler(rw http.ResponseWriter, req *http.Request) {
	upath := fmt.Sprintf("http://%s%s\n", rs.Addr, req.URL.Path)
	realIP := fmt.Sprintf("RemoteAddr=%s,X-Forwarded-For=%v,X-Real-Ip=%v\n", req.RemoteAddr, req.Header.Get("X-Forwarded-For"), req.Header.Get("X-Real-Ip"))
	header := fmt.Sprintf("headers =%v\n", req.Header)
	io.WriteString(rw, upath)
	io.WriteString(rw, realIP)
	io.WriteString(rw, header)
}

func (res *RealServer) ErrorHandler(rw http.ResponseWriter, req *http.Request) {
	upath := "error handler"
	rw.WriteHeader(500)
	io.WriteString(rw, upath)
}

func (res *RealServer) TimeOutHandler(rw http.ResponseWriter, req *http.Request) {
	time.Sleep(6 * time.Second)
	upath := "timeout handler"
	rw.WriteHeader(200)
	io.WriteString(rw, upath)
}
