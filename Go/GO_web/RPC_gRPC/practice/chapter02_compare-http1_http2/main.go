package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"time"
	"golang.org/x/net/http2"
)

func main() {

	// ---- HTTP1.1 CLient -----
	client1 := &http.Client{
		Transport: &http.Transport{	// Tại đây chúng ta có thể cấu hình các thông số liên quan đến HTTP1.1
			TLSNextProto: make(map[string]func(string, *tls.Conn) http.RoundTripper),	// Vô hiệu hóa HTTP2 bằng cách đặt TLSNextProto thành một map rỗng
		},
	}
	
	// ---- HTTP2 Client -----
	client2 := &http.Client{
		Transport: &http2.Transport{	
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // Cấu hình TLS để bỏ qua việc xác thực chứng chỉ, điều này thường được sử dụng trong môi trường phát triển hoặc khi làm việc với các máy chủ có chứng chỉ tự ký
		},
	}

	start := time.Now()
	for i := 0; i < 5; i++{
		resp, err := client1.Get("https://httpbin.org/get")
		if err == nil {
			resp.Body.Close()
		}
		fmt.Printf("Http/1.1 request %d completed\n", i+1)
	}

	fmt.Printf("Http/1.1 (5 requests): %v\n", time.Since(start))

	start = time.Now()
	for i := 0; i < 5; i++ {
		resp, err := client2.Get("https://httpbin.org/get")
		if err == nil {
			resp.Body.Close()
		}
		fmt.Printf("Http/2 request %d completed\n", i+1)
	}
	fmt.Printf("Http/2 (5 requests): %v\n", time.Since(start))
}