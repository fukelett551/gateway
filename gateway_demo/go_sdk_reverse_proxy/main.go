package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

const (
	server_addr = "http://127.0.0.1:8080/v1"
	proxy_port  = "8081"
)

//服务器 127.0.0.1:8080/v1/todo

//请求 127.0.0.1:8081/todo
//代理 127.0.0.1:8080/v1

// == 127.0.0.1:8080/v1/todo
//自动进行拼接
func main() {
	var (
		url *url.URL
		err error
	)
	if url, err = url.Parse(server_addr); err != nil {
		fmt.Println("parse addr error")
	}

	proxy := httputil.NewSingleHostReverseProxy(url)
	fmt.Println("proxy work on " + proxy_port)
	http.ListenAndServe(":"+proxy_port, proxy)
}
