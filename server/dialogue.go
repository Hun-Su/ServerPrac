package server

import (
	"echo/redis"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Dialogue struct {
	Id              string `json:"id"`
	Priority        string `json:"priority"`
	Sequence        string `json:"sequence"`
	Type            string `json:"type"`
	DirectionCamera string `json:"directionCamera"`
	TalkerImage     string `json:"talkerImage"`
	PlayCinematic   string `json:"playCinematic"`
	TextKor         string `json:"textKor"`
	TextEng         string `json:"textEng"`
	TextChn         string `json:"textChn"`
}

//20220531 leehs Redis 데이터를 가져와 저장
func (this Dialogue) Init(w http.ResponseWriter, req *http.Request) []Dialogue {
	tmp := redis.GetValue(cli, "dialogue")
	res := tmp.Val()
	dia := []Dialogue{}
	json.Unmarshal([]byte(res), &dia)

	return dia
}

//20220531 leehs id로 값 찾기
func (this Dialogue) GetDataByID(w http.ResponseWriter, req *http.Request) *Dialogue {
	dia := this.Init(w, req)
	id := req.FormValue("id")

	for _, i := range dia {
		//fmt.Println(i)
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
func (this Dialogue) GetDataByType(w http.ResponseWriter, req *http.Request) []Dialogue {
	dia := this.Init(w, req)
	Type := req.FormValue("type")
	var tmp []Dialogue

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
