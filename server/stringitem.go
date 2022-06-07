package server

import (
	"echo/redis"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type StringItem struct {
	Id      string `json:"id"`
	TextKor string `json:"textKor"`
	TextEng string `json:"textEng"`
	TextChn string `json:"textChn"`
}

//20220531 leehs Redis 데이터를 가져와 저장
func (this StringItem) Init(w http.ResponseWriter, req *http.Request) []StringItem {
	tmp := redis.GetValue(cli, "stringitem")
	res := tmp.Val()
	sq := []StringItem{}
	json.Unmarshal([]byte(res), &sq)

	return sq
}

//20220531 leehs id로 값 찾기
func (this StringItem) GetDataByID(w http.ResponseWriter, req *http.Request) *StringItem {
	si := this.Init(w, req)
	id := req.FormValue("id")

	for _, i := range si {
		if i.Id == id {
			tmp, _ := json.MarshalIndent(i, "", " ")
			fmt.Println(string(tmp))
			return &i
		}
	}
	log.Println("ID not found")
	return nil
}
