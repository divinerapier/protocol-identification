package other

import (
	"fmt"
	"log"
	"net/http"
)

// 包含 key = DAAP-Server
// 或者 value 包含 daap

func SendDAAP(addr string) {
	resp, err := http.Get("http://" + addr + "/server-info")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer resp.Body.Close()
	fmt.Println(resp.Header)
}
