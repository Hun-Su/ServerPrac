//package main
//
//import (
//	"fmt"
//	"log"
//	"net"
//)
//
//func main() {
//	conn, err := net.Dial("tcp", ":1234")
//	// "tcp", "tcp4", "tcp6", "unix" or "unixpacket". 로 설정 가능
//
//	if err != nil {
//		log.Fatalf("Failed to bind port ")
//	}
//
//	go StartReceive(conn)
//
//	for {
//		var s string
//		fmt.Scanln(&s)
//		conn.Write([]byte(s))
//	}
//}
//
//func StartReceive(conn net.Conn) {
//	recvBuf := make([]byte, 8192)
//
//	for {
//		n, err := conn.Read(recvBuf)
//
//		if err != nil || n <= 0 {
//			log.Println(err)
//			return
//		}
//
//		log.Println("From server : ", string(recvBuf[:n]))
//	}
//}
