package other

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func GrabGIT(addr string) {
	// transCfg := &http.Transport{
	// 	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	// }
	client := http.Client{
	//Transport: transCfg,
	}
	resp, err := client.Get("http://" + addr + "/text.git/HEAD")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(resp.Header)
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}
