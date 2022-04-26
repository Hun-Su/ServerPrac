package main

import (
	"echo/Config"
	"echo/HTTP_server"
	"echo/TCP_server"
	"echo/cli"
	"echo/db"
	"fmt"
	"log"
	"net"
	"net/http"
)

func main() {
	config := Config.LoadConfig()
	fmt.Println(config.Program)

	db := db.Get_Db()

	_, err := db.Query("SELECT name FROM emp where id = ?", 003)
	if err != nil {
		log.Fatal(err)
	}
	//
	//var name string
	//
	//for rows.Next() {
	//	err := rows.Scan(&name)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	fmt.Println(name)
	//}

	switch config.Program {
	case "HTTP":
		http.Handle("/", new(HTTP.TestHandler))

		http.ListenAndServe(config.Port.HTTP, nil)
	case "TCP":
		server, err := net.Listen("tcp", config.Port.TCP)
		if err != nil {
			log.Fatalln(err)
		}
		defer server.Close()

		log.Println("Server is running on:", config.Port.TCP)

		for {
			conn, err := server.Accept()
			go TCP_server.Comm(conn, err)
		}
	case "Client":
		conn, err := net.Dial("tcp", config.Port.TCP)
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
