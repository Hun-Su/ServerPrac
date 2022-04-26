package TCP_server

import (
	"bytes"
	"echo/Config"
	"io/ioutil"
	"log"
	"net"
	"net/http"
)

func Comm(c2 net.Conn, err error) {
	config := Config.LoadConfig()
	for {
		if err != nil {
			log.Println("Failed to accept conn.", err)
			continue
		}

		var c1 = c2

		buff := make([]byte, 8192)
		n, _ := c1.Read(buff)
		tmp1 := bytes.NewBuffer(buff[:n])

		println(tmp1)
		resp, err := http.Post("http://"+config.Port.HTTP, "text/plain", tmp1)

		if err != nil {
			panic(err)
		}

		body, err1 := ioutil.ReadAll(resp.Body)

		if err1 != nil {
			panic(err1)
		}

		c1.Write(body)

		log.Printf(string(body))
	}
}
