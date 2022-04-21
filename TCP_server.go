package main

//func comm(c2 net.Conn, err error) {
//	for {
//		if err != nil {
//			log.Println("Failed to accept conn.", err)
//			continue
//		}
//
//		var c1 = c2
//
//		buff := make([]byte, 8192)
//		n, _ := c1.Read(buff)
//		tmp1 := bytes.NewBuffer(buff[:n])
//
//		println(tmp1)
//		resp, err := http.Post("http://127.0.0.1:5000", "text/plain", tmp1)
//
//		if err != nil {
//			panic(err)
//		}
//
//		body, err1 := ioutil.ReadAll(resp.Body)
//
//		if err1 != nil {
//			panic(err1)
//		}
//
//		c1.Write(body)
//
//		log.Printf(string(body))
//	}
//}
//
//func main() {
//	addr := ":1234"
//	server, err := net.Listen("tcp", addr)
//	if err != nil {
//		log.Fatalln(err)
//	}
//	defer server.Close()
//
//	log.Println("Server is running on:", addr)
//
//	for {
//		conn, err := server.Accept()
//		go comm(conn, err)
//	}
//}
