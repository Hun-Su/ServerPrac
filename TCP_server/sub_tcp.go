package TCP_server

import (
	"bytes"
	"echo/logging"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
)

//leehs 20220623 서브 TCP 서버
func TCPServer() {
	listener, err := net.Listen("tcp", CONFIG.Port.SubTCP)
	if err != nil {
		logging.LogInfo(err.Error())
	}
	for {
		tcpconn, err := listener.Accept()
		if err != nil {
			logging.LogInfo(err.Error())
		}
		go TCPRead(tcpconn)
	}
}

//leehs 20220623 주어진 커넥션의 메시지를 읽어 출력
func TCPRead(tcpconn net.Conn) {
	buff := make([]byte, 8192)
	for {
		n, err := tcpconn.Read(buff)
		tmp1 := bytes.NewBuffer(buff[:n])
		if err != nil {
			logging.LogInfo(err.Error())
		}
		fmt.Println("subTCP:", tmp1)
		resp, err := http.Post("http://"+CONFIG.Port.HTTP, "text/plain", tmp1)

		if err != nil {
			logging.LogFatal(err.Error())
		}

		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			logging.LogFatal(err.Error())
		}
		fmt.Println(string(body))

		tcpconn.Write(body)
	}
}
