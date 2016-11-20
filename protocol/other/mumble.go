package other

import (
	"net"

	log "qiniupkg.com/x/log.v7"
)

// MumbleClient 。。。
type MumbleClient struct {
	tcpConn net.Conn
	udpConn net.Conn
	address string
}

// NewMumbleClient ...
func NewMumbleClient(addr string) *MumbleClient {
	return &MumbleClient{
		address: addr,
	}
}

// ConnectMumble ...
func (mc *MumbleClient) ConnectMumble() {
	tcpcon, err := net.Dial("tcp", mc.address)
	if err != nil {
		log.Fatalf("%v\n", err.Error())
	}
	mc.tcpConn = tcpcon
	udpcon, err := net.Dial("udp", mc.address)
	if err != nil {
		log.Fatalf("%v\n", err.Error())
	}
	mc.udpConn = udpcon
}
