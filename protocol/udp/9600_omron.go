package udp

import (
	"encoding/json"
	"fmt"
	"net"
	"time"
)

type OomronReply struct {
	ResponesCode      uint16
	ControllerModel   string
	ControllerVersion string
	ProgramAreaSize   uint16
	IOMSize           uint8
	NoOfDMWords       uint16
	TimeCounterSize   uint8
	ExpansionDMSize   uint8
	MemoryCard        uint8
	MemoryCardSize    uint16
}

func GrabOMRON(addr string) {
	dialer := net.Dialer{
		Timeout:  time.Second * 5,
		Deadline: time.Now().Add(time.Second * 5),
	}
	conn, err := dialer.Dial("udp", addr+":9600")
	if err != nil {
		fmt.Println(err)
		return
	}
	//conn.Write([]byte("\x80\x00\x02\x00\x00\x00\x00\x00\x00\x7a\x01\x01\x00\xcc\xcc\xcc\x00\x01"))

	conn.Write([]byte(
		"\x80\x00\x02\x00\x00\x00\x00\x63\x00\xef\x05\x01\x00"))

	var buf [1024]byte
	cnt, err := conn.Read(buf[:])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%v\n%s\n%x\n", buf[:cnt], buf[:cnt], buf[:cnt])

	resp := OomronReply{}

	resp.ResponesCode = uint16(buf[12])<<8 | uint16(buf[13])
	resp.ControllerModel = string(buf[14:34])
	resp.ControllerVersion = string(buf[34:55])
	resp.ProgramAreaSize = uint16(buf[94])<<8 | uint16(buf[95])
	resp.IOMSize = buf[96]
	resp.NoOfDMWords = uint16(buf[97])<<8 | uint16(buf[98])
	resp.TimeCounterSize = buf[99]
	resp.ExpansionDMSize = buf[100]
	resp.MemoryCard = buf[103]
	resp.MemoryCardSize = uint16(buf[104])<<8 | uint16(buf[105])
	data, _ := json.Marshal(&resp)
	fmt.Printf("%s\n", data)
}
