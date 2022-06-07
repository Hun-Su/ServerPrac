package server

import (
	"echo/redis"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type NPC struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	ShapeInfo string `json:"shapeInfo"`
	ShopType  string `json:"shopType"`
	Type      string `json:"type"`
}

//20220531 leehs Redis 데이터를 가져와 저장
func (this NPC) Init(w http.ResponseWriter, req *http.Request) []NPC {
	tmp := redis.GetValue(cli, "npc")
	res := tmp.Val()
	NPC := []NPC{}
	json.Unmarshal([]byte(res), &NPC)

	return NPC
}

//20220531 leehs id로 값 찾기
func (this NPC) GetDataByID(w http.ResponseWriter, req *http.Request) *NPC {
	npc := this.Init(w, req)
	id := req.FormValue("id")

	for _, i := range npc {
		if i.Id == id {
			tmp, _ := json.MarshalIndent(i, "", " ")
			fmt.Println(string(tmp))
			return &i
		}
	}
	log.Println("ID not found")
	return nil
}

//20220531 leehs Type으로 값 찾기
func (this NPC) GetDataByType(w http.ResponseWriter, req *http.Request) []NPC {
	dia := this.Init(w, req)
	Type := req.FormValue("type")
	var tmp []NPC

	for _, i := range dia {
		//fmt.Println(i)
		if i.Type == Type {
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
