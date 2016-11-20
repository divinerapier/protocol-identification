package udp

import (
	"fmt"
	"net"
	"time"
)

func GrabPCAnywhere(addr string) {
	dialer := net.Dialer{
		Timeout:  time.Second * 3,
		Deadline: time.Now().Add(time.Second * 3),
	}
	conn, err := dialer.Dial("udp", addr+":5632")
	if err != nil {
		fmt.Println(err)
		return
	}

	conn.Write([]byte("\x00\x00\x00\x00\x00\x00\x00\x00"))
	var buf [1024]byte
	cnt, err := conn.Read(buf[:])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("read %d\n", cnt)
	fmt.Printf("%s\n%v\n", buf[:cnt], buf[:cnt])
}
