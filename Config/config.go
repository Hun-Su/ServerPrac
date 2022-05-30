package Config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type TypeSelector struct {
	Program string `json:"program"`
	Port    struct {
		TCP  string `json:"tcp"`
		HTTP string `json:"http"`
	} `json:"port"`
	Db struct {
		User    string `json:"user"`
		Pwd     string `json:"pwd"`
		Port    string `json:"port"`
		Name    string `json:"name"`
		Timeout string `json:"timeout"`
	} `json:"db"`
	Redis struct {
		Port    string `json:"port"`
		Timeout string `json:"timeout"`
	} `json:"redis"`
}

func LoadConfig() TypeSelector {
	var config TypeSelector
	file, err := ioutil.ReadFile("pro.json")
	//defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	_ = json.Unmarshal(file, &config)
	return config
}
