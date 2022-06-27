package Client

import (
	"crypto/aes"
	"echo/crypto"
	"echo/logging"
	"log"
	"net"
)

func StartReceive(conn net.Conn) {
	recvBuf := make([]byte, 8192)

	for {
		n, err := conn.Read(recvBuf)

		if err != nil || n <= 0 {
			logging.LogInfo(err.Error())
		}
		key := "Hello Key 123456"
		block, err := aes.NewCipher([]byte(key))
		if err != nil {
			logging.LogInfo(err.Error())
			logging.Logger.Info(err.Error())
		}

		log.Println("From server : ", string(crypto.Decrypt(block, recvBuf[:n])))
	}
}
