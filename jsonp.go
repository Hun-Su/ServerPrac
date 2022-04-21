//package main
//
//import (
//	"encoding/json"
//	"fmt"
//	"log"
//	"os"
//)
//
//type MssqlConnect struct {
//	Program struct {
//		Type   string `json:"type"`
//		Server string `json:"server"`
//	} `json:"program"`
//}
//
//func LoadConfig() MssqlConnect {
//	var config MssqlConnect
//	file, err := os.Open("pro.json")
//	defer file.Close()
//	if err != nil {
//		log.Fatal(err)
//	}
//	decoder := json.NewDecoder(file)
//	err = decoder.Decode(&config)
//	return config
//}
//
//func main() {
//	config := LoadConfig()
//	fmt.Println(config.Program.Type)
//}
