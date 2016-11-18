package main

import (
	"protocol-identification/protocol/tcp"
)

func main() {
	// fmt.Println(protocol.CheckSupermicro("http://68.32.228.186:49152"))
	// fmt.Println(protocol.CheckSupermicro("http://58.226.233.207:49152"))
	// fmt.Println(protocol.CheckSupermicro("http://93.110.109.231:49152"))
	// fmt.Println(protocol.CheckSupermicro("http://59.54.120.89:49152"))
	// fmt.Println(protocol.CheckSupermicro("http://210.57.241.52:49152"))
	// fmt.Println(protocol.CheckSupermicro("http://39.125.234.197:49152"))

	//protocol.PPTPSend("5.148.22.104:1723")

	// protocol.IEC5104Send("221.156.99.103:2404")
	// protocol.IEC5104Send("41.223.105.53:2404")

	//protocol.SendDAAP("123.192.60.96:3689")

	// protocol.SendSIP("120.14.163.37")
	// protocol.SendSIP("79.122.34.174:5060")
	// protocol.SendSIP("84.187.201.180:5060")
	// protocol.SendSIP("128.234.103.143:5060")
	// protocol.SendSIP("85.240.166.90:5060")

	// protocol.SendVentrilo("174.136.101.106:3784")
	// protocol.SendVentrilo("220.233.185.152:3784")
	// protocol.SendVentrilo("142.222.67.10:3784")

	// protocol.SendGIT("85.25.185.118:9418")
	// protocol.SendGIT("77.88.23.115:9418")
	// protocol.SendGIT("54.183.13.26:9418")
	// protocol.SendGIT("45.79.218.81:9418")
	// protocol.SendGIT("85.25.185.118:9418")

	// ips := []string{
	// 	"166.143.158.234", "217.71.39.203", "217.201.245.135", "210.242.38.169", "31.61.112.192",
	// }
	// for _, v := range ips {
	// 	protocol.SendProconos(v + ":20547")
	// }

	// ips := []string{
	// 	"61.101.82.116",
	// }
	// for _, v := range ips {
	// 	protocol.SendHifly(v)
	// }

	// ips := []string{
	// 	"42.51.22.92",
	// }
	// for _, v := range ips {
	// 	protocol.SendFilezilla(v)
	// }
	// ips := []string{
	// 	"103.40.154.121",
	// 	"50.17.108.22",
	// 	"137.116.128.147",
	// }
	// for _, v := range ips {
	// 	fmt.Println(tcp.GrabAmqp(v))
	// }

	// ips := []string{
	// 	"184.4.107.14",
	// 	"99.11.134.131",
	// 	"104.53.177.12",
	// }
	// for _, v := range ips {
	// 	udp.GrabCoap(v)
	// }

	ips := []string{
		"10.10.10.234",
		"27.124.246.217",
		"119.56.147.83",
		"66.168.98.102",
		"189.144.162.255",
		"202.118.216.3",
		"109.129.179.180",
		"201.140.170.50",
		"80.14.222.249",
	}

	for _, v := range ips {
		tcp.GrabX11(v)
	}
}
