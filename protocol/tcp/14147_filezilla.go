package tcp

import (
	"fmt"
	"net"
	"time"
)

func SendFilezilla(addr string) {
	conn, err := net.DialTimeout("tcp", addr+":14147", time.Second*10)
	if err != nil {
		fmt.Println(err)
		return
	}
	conn.Write([]byte("AUTH TLS\r\n"))
	buf := make([]byte, 1024)
	cnt, _ := conn.Read(buf)
	fmt.Printf("count: %d\n", cnt)
	fmt.Printf("%s\n", buf)

}
