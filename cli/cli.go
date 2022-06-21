package Client

import (
	"crypto/aes"
	"echo/crypto"
	"fmt"
	"log"
	"net"
)

func StartReceive(conn net.Conn) {
	recvBuf := make([]byte, 8192)

	for {
		n, err := conn.Read(recvBuf)

		if err != nil || n <= 0 {
			log.Println(err)
			return
		}
		key := "Hello Key 123456"
		block, err := aes.NewCipher([]byte(key))
		if err != nil {
			fmt.Println(err)
			return
		}

		log.Println("From server : ", string(crypto.Decrypt(block, recvBuf[:n])))
	}
}
