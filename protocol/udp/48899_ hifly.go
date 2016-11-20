package udp

import (
	"fmt"
	"net"
	"time"
)

func GrabHifly(addr string) {
	conn, err := net.DialTimeout("udp", addr+":48899", time.Second*5)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
	conn.Write([]byte("12345678"))
	buf := make([]byte, 1024)
	cnt, _ := conn.Read(buf)
	fmt.Printf("count: %d\n", cnt)
	fmt.Printf("%s\n", buf)
}
