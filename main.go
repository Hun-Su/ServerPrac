package main

import (
	"echo/HTTP_server"
	"echo/TCP_server"
	"echo/cli"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

type Type_selector struct {
	Program string `json:"program"`
}

func LoadConfig() Type_selector {
	var config Type_selector
	file, err := os.Open("pro.json")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	return config
}

func main() {
	config := LoadConfig()
	fmt.Println(config.Program)

	switch config.Program {
	case "HTTP":
		http.Handle("/", new(HTTP.TestHandler))

		http.ListenAndServe("127.0.0.1:5000", nil)
	case "TCP":
		addr := ":1234"
		server, err := net.Listen("tcp", addr)
		if err != nil {
			log.Fatalln(err)
		}
		defer server.Close()

		log.Println("Server is running on:", addr)

		for {
			conn, err := server.Accept()
			go TCP_server.Comm(conn, err)
		}
	case "Client":
		conn, err := net.Dial("tcp", ":1234")
		// "tcp", "tcp4", "tcp6", "unix" or "unixpacket". 로 설정 가능

		if err != nil {
			log.Fatalf("Failed to bind port ")
		}

		go Client.StartReceive(conn)

		for {
			var s string
			fmt.Scanln(&s)
			conn.Write([]byte(s))
		}

	default:
		log.Println("Invalid Option")
		return
	}
}
