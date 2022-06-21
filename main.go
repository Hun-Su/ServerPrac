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
)

var CONFIG = Config.LoadConfig()

func main() {
	switch CONFIG.Program {
	case "HTTP":
		http.Handle("/", new(HTTP.TestHandler))
		http.ListenAndServe(CONFIG.Port.HTTP, nil)
	case "TCP":
		server, err := net.Listen("tcp", CONFIG.Port.TCP)
		if err != nil {
			log.Fatalln(err)
		}
		defer server.Close()

		log.Println("Server is running on:", CONFIG.Port.TCP)

		for {
			conn, _ := server.Accept()
			var handler = TCP_server.TcpSessionHandler(conn, conn.RemoteAddr())
			TCP_server.CliList = append(TCP_server.CliList, *handler)
			fmt.Println(TCP_server.CliList)
			go TCP_server.Comm(conn, err)
		}
	case "Client":
		conn, err := net.Dial("tcp", CONFIG.Port.TCP)
		if err != nil {
			log.Fatalf("Failed to bind port ")
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
	default:
		log.Println("Invalid Option")
		return
	}
}
