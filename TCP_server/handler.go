package TCP_server

import "net"

//20220620 leehs 클라이언트의 메인 tcp 서버와의 연결 정보와 sub tcp 서버와의 연결 정보 관리
type TcpHandler struct {
	Conn    net.Conn
	Key     net.Addr
	SubConn net.Conn
}

//20220620 leehs TcpHandler init
func TcpSessionHandler(conn net.Conn, net net.Addr, subconn net.Conn) (handler *TcpHandler) {
	handler = &TcpHandler{
		Conn:    conn,
		Key:     net,
		SubConn: subconn,
	}
	return
}
