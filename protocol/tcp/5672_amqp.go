package tcp

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"strings"
)

func GrabAmqp(addr string) bool {
	conn, err := net.Dial("tcp", addr+":5672")
	if err != nil {
		fmt.Println(err)
		return false
	}

	payload := new(bytes.Buffer)
	binary.Write(payload, binary.BigEndian, []byte("\x01\x00\x00"+
		"\x00\x00\x00\x05"+ // zise
		"\x00\x01\xf0\x00"+ // classid methodid
		"\x01\x00"))
	conn.Write(payload.Bytes())
	var buf [1024]byte

	cnt, err := conn.Read(buf[:])
	if err != nil {
		return false
	}
	return strings.Contains(string(buf[:cnt]), "AMQP")
}
