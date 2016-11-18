package tcp

import (
	"fmt"
	"net"
	"time"
)

func GrabVNC(addr string) {
	dialer := net.Dialer{
		Timeout:  time.Second * 5,
		Deadline: time.Now().Add(time.Second * 5),
	}

	conn, err := dialer.Dial("tcp", addr+":5900")
	if err != nil {
		fmt.Println(err)
		return
	}
	conn.Write([]byte("1234567890"))
	var buf [1024]byte
	cnt, err := conn.Read(buf[:])
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%v\n%s\n", buf[:cnt], buf[:cnt])

}
