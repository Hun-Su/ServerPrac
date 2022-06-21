package TCP_server

import "net"

type TcpHandler struct {
	Conn net.Conn
	Key  net.Addr
}

//20220620 leehs TcpHandler init
func TcpSessionHandler(conn net.Conn, net net.Addr) (handler *TcpHandler) {
	handler = &TcpHandler{
		Conn: conn,
		Key:  net,
	}
	return
}
