package Client

import (
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

		log.Println("From server : ", string(recvBuf[:n]))
	}
}
