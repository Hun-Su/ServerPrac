package server

import (
	"echo/redis"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Prop struct {
	Id             string `json:"id"`
	Description    string `json:"description"`
	HarvestType    string `json:"harvestType"`
	InreactTime    string `json:"inreactTime"`
	Name           string `json:"name"`
	RewardMaterial string `json:"rewardMaterial"`
	ShapeInfo      string `json:"shapeInfo"`
	Type           string `json:"type"`
}

//20220531 leehs Redis 데이터를 가져와 저장
func (this Prop) Init(w http.ResponseWriter, req *http.Request) []Prop {
	tmp := redis.GetValue(cli, "prop")
	res := tmp.Val()
	prop := []Prop{}
	json.Unmarshal([]byte(res), &prop)

	return prop
}

//20220531 leehs id로 값 찾기
func (this Prop) GetDataByID(w http.ResponseWriter, req *http.Request) *Prop {
	prop := this.Init(w, req)
	id := req.FormValue("id")

	for _, i := range prop {
		if i.Id == id {
			tmp, _ := json.MarshalIndent(i, "", " ")
			fmt.Println(string(tmp))
			return &i
		}
	}
	log.Println("ID not found")
	return nil
}

//20220531 leehs Name으로 값 찾기
func (this Prop) GetDataByName(w http.ResponseWriter, req *http.Request) []Prop {
	prop := this.Init(w, req)
	desc := req.FormValue("name")
	var tmp []Prop

	for _, i := range prop {
		if i.Name == desc {
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
