package main

import (
	"echo/Config"
	HTTP "echo/HTTP_server"
	"log"
	"net/http"
)

var CONFIG = Config.LoadConfig()

func main() {
	switch CONFIG.Program {
	case "HTTP":
		http.Handle("/", new(HTTP.TestHandler))
		http.ListenAndServe(CONFIG.Port.HTTP, nil)
	default:
		log.Println("Invalid Option")
		return
	}
}
