package server

import (
	"echo/redis"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Qitem struct {
	Id      string `json:"id"`
	TextKor string `json:"textKor"`
	TextEng string `json:"textEng"`
	TextChn string `json:"textChn"`
}

//20220531 leehs Redis 데이터를 가져와 저장
func (this Qitem) Init(w http.ResponseWriter, req *http.Request) []Qitem {
	tmp := redis.GetValue(cli, "qitem")
	res := tmp.Val()
	qitem := []Qitem{}
	json.Unmarshal([]byte(res), &qitem)

	return qitem
}

//20220531 leehs id로 값 찾기
func (this Qitem) GetDataByID(w http.ResponseWriter, req *http.Request) *Qitem {
	qitem := this.Init(w, req)
	id := req.FormValue("id")

	for _, i := range qitem {
		if i.Id == id {
			tmp, _ := json.MarshalIndent(i, "", " ")
			fmt.Println(string(tmp))
			return &i
		}
	}
	log.Println("ID not found")
	return nil
}
