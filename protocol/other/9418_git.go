package protocol

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func SendGIT(addr string) {
	// transCfg := &http.Transport{
	// 	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	// }
	client := http.Client{
	//Transport: transCfg,
	}
	resp, err := client.Get("http://" + addr + "/info/refs")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(resp.Header)
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}
