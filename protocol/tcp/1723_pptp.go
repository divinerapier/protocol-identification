package tcp

import (
	"fmt"
	"log"
	"net"
)

// PPTPSend ...
func PPTPSend(addr string) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatalln(err.Error())
	}

	// PPTPRequest 请求包
	// length           uint16
	// msgtype          uint16   // 1
	// magic            uint32   // 0x1A2B3C4D
	// controlmsgtype   uint16   // 1
	// reserved0        uint16   // 0
	// version          uint16   // 2
	// reserved1        uint16   // 0
	// framCapability   uint32   // 2
	// bearerCapability uint32   // 1
	// maxChannels      uint16   // 0
	// firmwareRevision uint16   // 1
	// hostName         [64]byte // dns host name
	// vendorString     [64]byte // 0

	conn.Write([]byte(
		"\x00\x9c" + // length 2bytes
			"\x00\x01" + // type 2bytes 1
			"\x1A\x2B\x3C\x4D" + // magic 4bytes 0x1A2B3C4D
			"\x00\x01" + // controlmsgtype 2bytes  1
			"\x00\x00" + //reserved0 2bytes 0
			"\x00\x02" + //version 2bytes 2
			"\x00\x00" + //reserved1 2bytes 0
			"\x00\x00\x00\x02" + //framCapability 4bytes 2
			"\x00\x00\x00\x01" + //bearerCapability 4bytes 1
			"\x00\x00" + //maxChannels 2bytes 0
			"\x00\x01" + //firmwareRevision 2bytes 1
			"00000000000000000000000000000000" + //hostName
			"00000000000000000000000000000000" + //hostName
			"00000000000000000000000000000000" + //vendorString
			"00000000000000000000000000000000", //vendorString

	))
	buf := make([]byte, 1024)
	cnt, _ := conn.Read(buf)
	fmt.Printf("read %d\n%s\n", cnt, buf)
	r := parseReply(buf)
	fmt.Println("\n\n")
	fmt.Println(r)
}

type reply struct {
	length     uint16
	msgtype    uint16
	magic      uint32
	cmt        uint16
	reserved0  uint16
	version    uint16
	resultCode uint8
	errCode    uint8
	framing    uint32
	bearer     uint32
	channels   uint16
	firmware   uint16
	host       string
	vendor     string
}

func parseReply(r []byte) *reply {
	length := uint16(r[0])<<8 + uint16(r[1])
	msgtype := uint16(r[2])<<8 + uint16(r[3])
	magic := uint32(r[4])<<24 + uint32(r[5])<<16 + uint32(r[6])<<8 + uint32(r[7])
	cmt := uint16(r[8])<<8 + uint16(r[9])
	reserved0 := uint16(r[10])<<8 + uint16(r[11])
	version := uint16(r[12])<<8 + uint16(r[13])
	resCode := uint8(r[14])
	errCode := uint8(r[15])
	framing := uint32(r[16])<<24 + uint32(r[17])<<16 + uint32(r[18])<<8 + uint32(r[19])
	bearer := uint32(r[20])<<24 + uint32(r[21])<<16 + uint32(r[22])<<8 + uint32(r[23])
	channels := uint16(r[24])<<8 + uint16(r[25])
	firmware := uint16(r[26])<<8 + uint16(r[27])
	return &reply{
		length:     length,
		msgtype:    msgtype,
		magic:      magic,
		cmt:        cmt,
		reserved0:  reserved0,
		version:    version,
		resultCode: resCode,
		errCode:    errCode,
		framing:    framing,
		bearer:     bearer,
		channels:   channels,
		firmware:   firmware,
		host:       string(r[28 : 28+64]),
		vendor:     string(r[28+64 : 28+64+64]),
	}
}
