package tcp

import (
	"fmt"
	"net"

	log "qiniupkg.com/x/log.v7"
)

/*

   0                   1                   2                   3
    0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
   |           启动字符 68H         |
   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
   |         APDU长度(最大253)      |
   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
   |          控制域八位位组 1       |   APCI
   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
   |          控制域八位位组 2       |
   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
   |          控制域八位位组 3       |
   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+                     APDU
   |          控制域八位位组 4       |
   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
   |                               |
   +                               +
   |                               |  ASDU
   |                               |
   +                               +
   |                               |
   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

*/

func IEC5104Send(addr string) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Errorf("%v\n", err.Error())
		return
	}
	conn.Write([]byte(
		"\x68" +
			"\x04" + // 长度
			"\x01" +
			"\x00\x00\x00"))

	buf := make([]byte, 1024)
	cnt, err := conn.Read(buf)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Printf("read %d\n", cnt)
	for i := 0; i < cnt; i++ {
		fmt.Printf("%x ", buf[i])
		if (i+1)%40 == 0 {
			fmt.Println()
		}
	}
	fmt.Printf("%s\n", buf)
}
