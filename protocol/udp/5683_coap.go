package udp

import (
	"fmt"
	"net"
	"time"
)

func GrabCoap(addr string) {
	var mapCode = make(map[int]string)
	mapCode[65] = "Created"
	mapCode[66] = "Deleted"
	mapCode[67] = "Valid"
	mapCode[68] = "Changed"
	mapCode[69] = "Content"
	mapCode[128] = "BadRequest"
	mapCode[129] = "Unauthorized"
	mapCode[130] = "BadOption"
	mapCode[131] = "Forbidden"
	mapCode[132] = "NotFound"
	mapCode[133] = "MethodNotAllowed"
	mapCode[134] = "NotAcceptable"
	mapCode[140] = "PreconditionFailed"
	mapCode[141] = "RequestEntityTooLarge"
	mapCode[143] = "UnsupportedMediaType"
	mapCode[160] = "InternalServerError"
	mapCode[161] = "NotImplemented"
	mapCode[162] = "BadGateway"
	mapCode[163] = "ServiceUnavailable"
	mapCode[164] = "GatewayTimeout"
	mapCode[165] = "ProxyingNotSupported"
	mapType := make(map[int]string, 4)
	mapType[0] = "Confirmable"
	mapType[1] = "NonConfirmable"
	mapType[2] = "Acknowledgement"
	mapType[3] = "Reset"

	conn, err := net.DialTimeout("udp", addr+":5683", time.Second*3)
	if err != nil {
		fmt.Println(err)
		return
	}
	conn.Write([]byte("\x40\x01\x7d\x35\xff1qaz2wsx3edc"))
	var buf [1024]byte
	_, err = conn.Read(buf[:])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("version: %d\n", buf[0]>>6)
	fmt.Printf("type: %s\n", mapType[int(buf[0]&0x30>>4)])
	fmt.Printf("token length: %d\n", buf[0]&0x0f)
	if a, ok := mapCode[int(buf[1])]; ok {
		fmt.Printf("code: %s\n", a)
	} else {
		fmt.Printf("code: unknown\n")
	}
	fmt.Printf("message id: %d\n", int(buf[2])<<8|int(buf[3]))
	fmt.Printf("token: 0x%x\n", int(buf[4])<<8|int(buf[5]))
}
