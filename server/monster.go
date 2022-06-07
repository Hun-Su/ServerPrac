package server

import (
	"echo/redis"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Monster struct {
	Id           string `json:"id"`
	AP           string `json:"ap"`
	DP           string `json:"dp"`
	HP           string `json:"hp"`
	Attitude     string `json:"attitude"`
	Chase        string `json:"chase"`
	Exp          string `json:"exp"`
	Grade        string `json:"grade"`
	HitRate      string `json:"hitRate"`
	HitSpeed     string `json:"hitSpeed"`
	Immune       string `json:"immune"`
	IsGuard      string `json:"isGuard"`
	Level        string `json:"level"`
	MovePattern  string `json:"movePattern"`
	Name         string `json:"name"`
	Relationship string `json:"relationship"`
	ShapeInfo    string `json:"shapeInfo"`
	Sight        string `json:"sight"`
	SkillPattern string `json:"skillPattern"`
	SpriteInfo   string `json:"spriteInfo"`
	Type         string `json:"type"`
	UseSkill     string `json:"useSkill"`
}

//20220531 leehs Redis 데이터를 가져와 저장
func (this Monster) Init(w http.ResponseWriter, req *http.Request) []Monster {
	tmp := redis.GetValue(cli, "monster")
	res := tmp.Val()
	monster := []Monster{}
	json.Unmarshal([]byte(res), &monster)

	return monster
}

//20220531 leehs id로 값 찾기
func (this Monster) GetDataByID(w http.ResponseWriter, req *http.Request) *Monster {
	monster := this.Init(w, req)
	id := req.FormValue("id")

	for _, i := range monster {
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
func (this Monster) GetDataByName(w http.ResponseWriter, req *http.Request) []Monster {
	monster := this.Init(w, req)
	desc := req.FormValue("name")
	var tmp []Monster

	for _, i := range monster {
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
