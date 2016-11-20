package other

import (
	"fmt"
	"net/http"
	"strings"
)

/*
HTTP/1.0 404 Not Found
SERVER: Linux/2.4.19-rmk4, UPnP/1.0, Intel SDK for UPnP devices /1.2
CONNECTION: close
CONTENT-LENGTH: 48
CONTENT-TYPE: text/html
*/

func CheckSupermicro(addr string) bool {
	resp, err := http.Get(addr)
	if err != nil {
		fmt.Printf("%v\n", err.Error())
		return false
	}
	if servers, ok := resp.Header["Server"]; ok {
		server := servers[0]
		if strings.Contains(server, "Linux/") && strings.Contains(server, "UPnP/") && (strings.Contains(server, "SDK for UPnP devices") || (strings.Contains(server, "Upnp SDK"))) {
			return true
		}
	}
	fmt.Println("server not found")
	return false
}
