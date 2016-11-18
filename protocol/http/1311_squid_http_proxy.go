package http

import (
	"fmt"
	"net/http"
	"strings"
)

/*

Squid http proxyVersion: 3.4.11
HTTP/1.1 400 Bad Request
Server: squid/3.4.11
Mime-Version: 1.0
Date: Sat, 05 Nov 2016 21:45:35 GMT
Content-Type: text/html
Content-Length: 3193
X-Squid-Error: ERR_INVALID_URL 0
Vary: Accept-Language
Content-Language: en
X-Cache: MISS from offenbach.actproxy.com
Via: 1.1 offenbach.actproxy.com (squid/3.4.11)
Connection: close

*/

func squid_server_grab(host, port string) bool {
	if host[:4] != "http" {
		host = "http://" + host
	}
	resp, err := http.Get(host + port)
	if err != nil {
		fmt.Printf("%v\n", err.Error())
		return false
	}
	servers, ok := resp.Header["Server"]
	if !ok || strings.Contains(servers[0], "squid") {
		fmt.Println("not squid server")
		return false
	}

	return true
}
