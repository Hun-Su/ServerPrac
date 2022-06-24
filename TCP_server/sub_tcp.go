package TCP_server

import (
	"bytes"
	"fmt"
	"log"
	"net"
)

//leehs 20220623 서브 TCP 서버
func TCPServer() {
	listener, err := net.Listen("tcp", CONFIG.Port.SubTCP)
	if err != nil {
		log.Println(err)
		return
	}
	for {
		tcpconn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go TCPRead(tcpconn)
	}
}

//leehs 20220623 주어진 커넥션의 메시지를 읽어 출력
func TCPRead(tcpconn net.Conn) {
	buff := make([]byte, 8192)
	for {
		n, error := tcpconn.Read(buff)
		tmp1 := bytes.NewBuffer(buff[:n])
		if error != nil {
			log.Println(error)
			return
		}
		fmt.Println("subTCP:", tmp1)
	}
}
