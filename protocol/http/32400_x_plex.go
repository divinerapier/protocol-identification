package http

import (
	"bytes"
	"fmt"
	"net"
	"time"
)

func GrabXPlex(addr string) {
	dialer := net.Dialer{
		Timeout:  time.Second * 5,
		Deadline: time.Now().Add(time.Second * 5),
	}

	conn, err := dialer.Dial("tcp", addr+":32400")
	if err != nil {
		fmt.Println(err)
		return
	}
	conn.Write([]byte("GET /0okmnji9 HTTP/1.1\r\nHost: " + addr + ":5985\r\nUser-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/54.0.2840.71 Safari/537.36\r\nAccept: */*\r\n\r\n"))
	var buf [1024]byte
	cnt, err := conn.Read(buf[:])
	if err != nil {
		fmt.Println(err)
		return
	}
	if i := bytes.Index(buf[:cnt], []byte("X-Plex-Protocol")); i > 0 {
		end := bytes.Index(buf[i:cnt], []byte("\r\n"))
		banner := buf[i : i+end]
		fmt.Printf("%s\n", banner)
	}
}
