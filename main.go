package main

import (
	"bufio"
	"crypto/aes"
	"echo/Config"
	HTTP "echo/HTTP_server"
	"echo/TCP_server"
	Client "echo/cli"
	"echo/crypto"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"sync"
)

var CONFIG = Config.LoadConfig()

func main() {
	var wg sync.WaitGroup
	switch CONFIG.Program {
	case "HTTP":
		serveHTTP()
	case "TCP":
		serveTCP()
	case "Client":
		serveClient()
	case "subTCP":
		servesubTCP()
	//leehs 20220623 4개의 역활을 동시 수행
	case "All":
		wg.Add(4)

		go serveClient()
		go serveTCP()
		wg.Done()
		go servesubTCP()
		wg.Done()
		go serveHTTP()
		wg.Done()

		wg.Wait()
	default:
		log.Println("Invalid Option")
		return
	}
}

//leehs 20220623 HTTP 서버
func serveHTTP() {
	http.Handle("/", new(HTTP.TestHandler))
	http.ListenAndServe(CONFIG.Port.HTTP, nil)
}

//leehs 20220623 메인 TCP 서버
func serveTCP() {
	server, err := net.Listen("tcp", CONFIG.Port.TCP)
	if err != nil {
		log.Fatalln(err)
	}
	defer server.Close()

	log.Println("Server is running on:", CONFIG.Port.TCP)

	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		tcpconn, err := net.Dial("tcp", CONFIG.Port.SubTCP)
		if err != nil {
			log.Fatalln(err)
		}

		//leehs 20220623 접속해 있는 클리이언트 리스트와 연결 정보
		var handler = TCP_server.TcpSessionHandler(conn, conn.RemoteAddr(), tcpconn)
		TCP_server.CliList = append(TCP_server.CliList, *handler)
		fmt.Println(TCP_server.CliList)

		//leehs 20220623 TCP 핸들러
		go TCP_server.Comm(handler)
	}
}

//leehs 20220623 sub TCP 서버
func servesubTCP() {
	tcpconn := TCP_server.TCPServer()
	TCP_server.TCPRead(tcpconn)
}

//leehs 20220623 Client
func serveClient() {
	conn, err := net.Dial("tcp", CONFIG.Port.TCP)
	if err != nil {
		log.Fatalf("Failed to bind port")
	}

	go Client.StartReceive(conn)

	key := "Hello Key 123456"
	for {
		block, err := aes.NewCipher([]byte(key))
		if err != nil {
			fmt.Println(err)
			return
		}
		r := bufio.NewReader(os.Stdin)
		s, _ := r.ReadString('\n')
		if strings.TrimSpace(s) == "quit" {
			log.Println("Closing connection")
			conn.Write([]byte("quit"))
			break
		}
		msg := crypto.Encrypt(block, []byte(s))
		fmt.Println(msg)

		conn.Write(msg)
	}
}
