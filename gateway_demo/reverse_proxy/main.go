package main

import (
	"bufio"
	"fmt"
	"net/http"
	"net/url"
)

const (
	server_addr = "http://127.0.0.1:8080"
	proxy_port  = "8081"
)

func handler(resp http.ResponseWriter, req *http.Request) {
	var (
		server    *url.URL
		err       error
		res       *http.Response
		transport http.RoundTripper
	)
	if server, err = url.Parse(server_addr); err != nil {
		fmt.Println("parse err")
	}
	req.URL.Scheme = server.Scheme
	req.URL.Host = server.Host
	fmt.Println(req.URL.Scheme)
	fmt.Println(req.URL.Host)

	transport = http.DefaultTransport
	if res, err = transport.RoundTrip(req); err != nil {
		fmt.Println("roundtrip err")
	}

	for key, value := range res.Header {
		for _, v := range value {
			resp.Header().Add(key, v)
		}
	}
	defer res.Body.Close()
	bufio.NewReader(res.Body).WriteTo(resp)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("proxy_server port work at " + proxy_port)
	http.ListenAndServe(":"+proxy_port, nil)
}
