package Config

import (
	"encoding/json"
	"log"
	"os"
)

type Type_selector struct {
	Program string `json:"program"`
	Port    struct {
		TCP  string `json:"tcp"`
		HTTP string `json:"http"`
	} `json:"port"`
	Db struct {
		User string `json:"user"`
		Pwd  string `json:"pwd"`
		Port string `json:"port"`
		Name string `json:"name"`
	} `json:"db"`
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
