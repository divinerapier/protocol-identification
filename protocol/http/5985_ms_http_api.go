package http

import (
	"bytes"
	"fmt"
	"net"
	"time"
)

func GrabMicrosoftHTTPAPI(addr string) {
	dialer := net.Dialer{
		Timeout:  time.Second * 3,
		Deadline: time.Now().Add(time.Second * 3),
	}
	conn, err := dialer.Dial("tcp", addr+":5985")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	conn.Write([]byte("GET /0okmnji9 HTTP/1.1\r\nHost: " + addr + ":5985\r\nUser-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/54.0.2840.71 Safari/537.36\r\nAccept: */*\r\n\r\n"))
	var buf [10240]byte
	cnt, err := conn.Read(buf[:])
	if err != nil {
		fmt.Println(err)
		return
	}
	if bytes.Contains(buf[:cnt], []byte("\r\n\r\n")) {
		if index := bytes.Index(buf[:cnt], []byte("Server")); index > 0 {
			end := bytes.Index(buf[index:cnt], []byte("\r\n"))
			banner := buf[index : index+end]
			fmt.Printf("banner: %s\n", banner)
		}
	}
}
