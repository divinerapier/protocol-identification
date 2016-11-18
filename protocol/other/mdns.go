package protocol

import "net"

func PingMDNS(ip string) (net.Conn, error) {
	return net.Dial("udp", ip+":5353")
}

func RecvMDNS(c *net.Conn, data []byte) ([]byte, error) {
	_, err := (*c).Read(data)
	return data, err
}

func SendMDNS(c *net.Conn, data []byte) error {
	_, err := (*c).Write(data)
	return err
}
