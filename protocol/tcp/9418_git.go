package tcp

import (
	"fmt"
	"net"
	"time"
)

func GrabGIT(addr string) {
	dialer := net.Dialer{
		Timeout:  time.Second * 10,
		Deadline: time.Now().Add(time.Second * 10),
	}
	conn, err := dialer.Dial("tcp", addr+":9418")
	if err != nil {
		fmt.Println(err)
		return
	}
	conn.Write([]byte("\x30\x30\x33\x62\x67\x69\x74\x2d\x75\x70\x6c\x6f\x61\x64\x2d\x70\x61\x63\x6b\x20\x2f\x31\x71\x61\x7a\x32\x77\x73\x78\x2f\x30\x6f\x6b\x6d\x39\x69\x6a\x6e\x2e\x67\x69\x74\x00\x68\x6f\x73\x74\x3d\x67\x69\x74\x68\x75\x62\x2e\x63\x6f\x6d\x00"))
	var buf [1024]byte
	cnt, err := conn.Read(buf[:])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s\n%v\n%x\n", buf[:cnt], buf[:cnt], buf[:cnt])

	// packetLength, err := strconv.ParseInt(string(buf[0:4]), 16, 0)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	packetLength := int64(buf[0]-'0')<<12 | int64(buf[1]-'0')<<8 | int64(buf[2]-'0')<<4 | int64(buf[3]-'0')
	if int64(cnt) != packetLength {
		fmt.Println("not git")
		return
	}
	fmt.Printf("cnt = %d\npacket_length = %d\n", cnt, packetLength)
}
