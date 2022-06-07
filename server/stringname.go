package server

import (
	"echo/redis"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type StringName struct {
	Id      string `json:"id"`
	TextKor string `json:"textKor"`
	TextEng string `json:"textEng"`
	TextChn string `json:"textChn"`
}

//20220531 leehs Redis 데이터를 가져와 저장
func (this StringName) Init(w http.ResponseWriter, req *http.Request) []StringName {
	tmp := redis.GetValue(cli, "stringname")
	res := tmp.Val()
	sn := []StringName{}
	json.Unmarshal([]byte(res), &sn)

	return sn
}

//20220531 leehs id로 값 찾기
func (this StringName) GetDataByID(w http.ResponseWriter, req *http.Request) *StringName {
	sn := this.Init(w, req)
	id := req.FormValue("id")

	for _, i := range sn {
		if i.Id == id {
			tmp, _ := json.MarshalIndent(i, "", " ")
			fmt.Println(string(tmp))
			return &i
		}
	}
	log.Println("ID not found")
	return nil
}
