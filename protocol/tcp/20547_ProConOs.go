package tcp

import (
	"bytes"
	"fmt"
	"net"
	"time"
)

func SendProconos(addr string) {
	keys := []string{"Ladder Logic Runtime", "PLC Type", "Project Name", "Boot Project", "Project Source Code"}
	m := make(map[string]string)
	conn, err := net.DialTimeout("tcp", addr, time.Second*3)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	conn.Write([]byte("\xcc\x01\x00\x0b\x40\x02\x00\x00\x47\xee"))
	buf := make([]byte, 1024)
	cnt, _ := conn.Read(buf)
	if buf[0] != '\xcc' {
		return
	}
	buf = buf[12:cnt]
	buf = bytes.Replace(buf, []byte("\x00\x00"), []byte("\x00"), -1)
	fields := bytes.Split(buf, []byte("\x00"))
	i := 0
	for _, v := range fields {
		if string(v) == "" {
			continue
		}
		m[keys[i]] = string(v)
		i++
		if i == 5 {
			break
		}
	}
	for k, v := range m {
		fmt.Printf("%s: %s\n", k, v)
	}
}
