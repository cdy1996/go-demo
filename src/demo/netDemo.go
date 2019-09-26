package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	//conn, e := net.Dial("tcp", "127.0.0.1")
	//conn.SetDeadline(time.Now().Add(1*time.Second))
	//conn.SetReadDeadline(time.Now().Add(3*time.Second))
	//conn.SetWriteDeadline(time.Now().Add(3*time.Second))

	httpServerTest()

	time.Sleep(5 * time.Minute)

}

func httpTest() {
	//client := http.DefaultClient
	//tripper := client.Transport
	url1 := "http://www.baidu.com"
	fmt.Printf("Send request to %q with method GET ...\n", url1)
	resp1, err := http.Get(url1)
	if err != nil {
		fmt.Printf("request sending error: %v\n", err)
	}
	defer resp1.Body.Close()
	line1 := resp1.Proto + " " + resp1.Status
	fmt.Printf("The first line of response:\n%s\n", line1)
}

func httpServerTest() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`hello world`))
	})
	http.ListenAndServe("127.0.0.1:8080", nil)
}
