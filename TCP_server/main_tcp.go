package TCP_server

import (
	"bytes"
	"crypto/aes"
	"echo/Config"
	"echo/crypto"
	"fmt"
	"log"
	"strings"
)

var CONFIG = Config.LoadConfig()

//20220620 leehs 클라이언트의 정보를 저장하는 리스트
var CliList []TcpHandler

//20220620 leehs tcp 서버 핸들러 클라이언트와의 연결 c1, 서브 TCP와의 연결 c2
func Comm(handler *TcpHandler) {
	fmt.Printf("Serving %s\n", handler.Conn.RemoteAddr().String())
	for {
		key := "Hello Key 123456"
		block, err := aes.NewCipher([]byte(key))

		//20220620 leehs 클라이언트의 메시지를 읽은 후 quit일 경우 해당 클라이언트 정보 삭제
		//ctl+c로 연결 해제시 다시 없앨 수 없는 문제
		buff := make([]byte, 8192)
		n, _ := handler.Conn.Read(buff)
		tmp1 := bytes.NewBuffer(buff[:n])

		if err != nil {
			log.Println(err)
			return
		}
		log.Println("From ", handler.Conn.RemoteAddr(), ": ", tmp1)

		if strings.TrimSpace(tmp1.String()) == "quit" || handler.Conn.RemoteAddr() == nil {
			for i, j := range CliList {
				if j.Key == handler.Conn.RemoteAddr() {
					fmt.Println("Client ", j.Key, " Disconnected")
					CliList = append(CliList[:i], CliList[i+1:]...)
					fmt.Println(CliList)
					break
				}
			}
			return
		}

		if tmp1.String() == "" {
			return
		}
		body := crypto.Decrypt(block, tmp1.Bytes())
		handler.SubConn.Write(body)
		//20220620 leehs client에게 다시 보내기 위해 재암호화
		res := crypto.Encrypt(block, body)
		handler.Conn.Write(res)
		log.Println("From ", handler.Conn.RemoteAddr(), ": ", string(body))
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
	}
}
