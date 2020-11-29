package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
)

type Proxy struct{}

func (p *Proxy) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	fmt.Println("receive request ", req.Method, req.Host, req.RemoteAddr)
	transport := http.DefaultTransport
	//浅拷贝请求
	outReq := new(http.Request)
	*outReq = *req
	//设置http头部，并向下游发送请求
	if clientIp, _, err := net.SplitHostPort(req.RemoteAddr); err != nil {
		fmt.Println("err", err.Error())
		return
	} else {
		if prior, ok := outReq.Header["X-Forworded-For"]; ok {
			clientIp = strings.Join(prior, ",") + "," + clientIp
		}
		outReq.Header.Set("X-Forworded-For", clientIp)
	}

	res, err := transport.RoundTrip(outReq)
	if err != nil {
		resp.WriteHeader(http.StatusBadGateway)
		return
	}
	//回写http
	for key, value := range res.Header {
		for _, v := range value {
			resp.Header().Add(key, v)
		}
	}
	resp.WriteHeader(res.StatusCode)
	io.Copy(resp, res.Body)
	res.Body.Close()

}

func main() {
	fmt.Println("server on 8080")
	http.Handle("/", &Proxy{})
	http.ListenAndServe("0.0.0.0:8080", nil)
}
