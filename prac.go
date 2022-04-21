package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"unicode"
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
		HTTP_server()
	case "TCP":
		TCP_server()
	case "Client":
		client()
	default:
		log.Println("Invalid Option")
		return
	}
}

func HTTP_server() {
	http.Handle("/", new(testHandler))

	http.ListenAndServe("127.0.0.1:5000", nil)
}

type testHandler struct {
	http.Handler
}

func (h *testHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	b, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()

	var operator string
	var a []string
	var res string

	for _, val := range string(b) {
		if !unicode.IsDigit(val) {
			operator = string(val)
		}
	}

	a = strings.Split(string(b), operator)

	fmt.Println(a)
	v1, _ := strconv.Atoi(a[0])
	v2, _ := strconv.Atoi(a[1])

	switch operator {
	case "+":
		res = strconv.Itoa(v1 + v2)
	case "-":
		res = strconv.Itoa(v1 - v2)
	case "*":
		res = strconv.Itoa(v1 * v2)
	case "/":
		res = strconv.Itoa(v1 / v2)
	default:
		res = "Not a binary equation"
	}

	fmt.Println(res)
	w.Write([]byte(res))
}

func comm(c2 net.Conn, err error) {
	for {
		if err != nil {
			log.Println("Failed to accept conn.", err)
			continue
		}

		var c1 = c2

		buff := make([]byte, 8192)
		n, _ := c1.Read(buff)
		tmp1 := bytes.NewBuffer(buff[:n])

		println(tmp1)
		resp, err := http.Post("http://127.0.0.1:5000", "text/plain", tmp1)

		if err != nil {
			panic(err)
		}

		body, err1 := ioutil.ReadAll(resp.Body)

		if err1 != nil {
			panic(err1)
		}

		c1.Write(body)

		log.Printf(string(body))
	}
}

func TCP_server() {
	addr := ":1234"
	server, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln(err)
	}
	defer server.Close()

	log.Println("Server is running on:", addr)

	for {
		conn, err := server.Accept()
		go comm(conn, err)
	}
}

func client() {
	conn, err := net.Dial("tcp", ":1234")
	// "tcp", "tcp4", "tcp6", "unix" or "unixpacket". 로 설정 가능

	if err != nil {
		log.Fatalf("Failed to bind port ")
	}

	go StartReceive(conn)

	for {
		var s string
		fmt.Scanln(&s)
		conn.Write([]byte(s))
	}
}

func StartReceive(conn net.Conn) {
	recvBuf := make([]byte, 8192)

	for {
		n, err := conn.Read(recvBuf)

		if err != nil || n <= 0 {
			log.Println(err)
			return
		}

		log.Println("From server : ", string(recvBuf[:n]))
	}
}
