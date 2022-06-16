package TCP_server

import (
	"bytes"
	"crypto/aes"
	"echo/Config"
	"echo/crypto"
	"fmt"
	"log"
	"net"
)

var CONFIG = Config.LoadConfig()

func Comm(c2 net.Conn, err error) {
	for {
		key := "Hello Key 123456"
		if err != nil {
			log.Println("Failed to accept conn.", err)
			continue
		}

		var c1 = c2

		buff := make([]byte, 8192)
		n, _ := c1.Read(buff)
		tmp1 := bytes.NewBuffer(buff[:n])
		fmt.Println(tmp1)

		block, err := aes.NewCipher([]byte(key))
		if err != nil {
			fmt.Println(err)
			return
		}
		//resp, err := http.Post("http://"+config.Port.HTTP, "text/plain", tmp1)
		//
		//if err != nil {
		//	panic(err)
		//}
		//
		//body, err1 := ioutil.ReadAll(resp.Body)
		//
		//if err1 != nil {
		//	panic(err1)
		//}
		body := crypto.Decrypt(block, tmp1.Bytes())
		c1.Write(body)

		log.Printf(string(body))
	}
}
