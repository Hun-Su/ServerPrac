package server

import (
	"echo/redis"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type StringQuest struct {
	Id          string `json:"id"`
	Description string `json:"description"`
}

//20220531 leehs Redis 데이터를 가져와 저장
func (this StringQuest) Init(w http.ResponseWriter, req *http.Request) []StringQuest {
	tmp := redis.GetValue(cli, "stringquest")
	res := tmp.Val()
	sq := []StringQuest{}
	json.Unmarshal([]byte(res), &sq)

	return sq
}

//20220531 leehs id로 값 찾기
func (this StringQuest) GetDataByID(w http.ResponseWriter, req *http.Request) *StringQuest {
	sq := this.Init(w, req)
	id := req.FormValue("id")

	for _, i := range sq {
		if i.Id == id {
			tmp, _ := json.MarshalIndent(i, "", " ")
			fmt.Println(string(tmp))
			return &i
		}
	}
	log.Println("ID not found")
	return nil
}

//20220531 leehs Desciption으로 값 찾기
func (this StringQuest) GetDataByDesc(w http.ResponseWriter, req *http.Request) []StringQuest {
	sq := this.Init(w, req)
	desc := req.FormValue("desc")
	var tmp []StringQuest

	for _, i := range sq {
		if i.Description == desc {
			tmp = append(tmp, i)
		}
	}
	for _, i := range tmp {
		tmp1, _ := json.MarshalIndent(i, "", " ")
		fmt.Println(string(tmp1))
	}
	if len(tmp) == 0 {
		log.Println("ID not found")
	}
	return tmp
}
