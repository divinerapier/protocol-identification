package other

import (
	"net"

	log "qiniupkg.com/x/log.v7"
)

// ConnectDahua ...
func ConnectDahua(ip string) *net.Conn {
	conn, err := net.Dial("tcp", ip+":37777")
	if err != nil {
		log.Errorf("%v\n", err.Error())
	}
	return &conn
}

// SendDahua ...
func SendDahua(c *net.Conn, msg []byte) error {
	_, err := (*c).Write(msg)
	return err
}

// RecvDahua ...
func RecvDahua(c *net.Conn, data []byte) ([]byte, error) {
	_, err := (*c).Read(data)
	return data, err
}
