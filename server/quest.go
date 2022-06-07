package server

import (
	"echo/redis"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Quest struct {
	Id                   string `json:"id"`
	Description          string `json:"description"`
	DialogueAfterAccept  string `json:"dialogueAfterAccept"`
	DialogueBeforeAccept string `json:"dialogueBeforeAccept"`
	DialogueClear        string `json:"dialogueClear"`
	ConditionId          string `json:"conditionId"`
	Name                 string `json:"name"`
	NextQuest            string `json:"nextQuest"`
	RewardItemId         string `json:"rewardItemId"`
	SubjectName          string `json:"subjectName"`
	SubjectType          string `json:"subjectType"`
	SubjectCoundId       string `json:"subjectCoundId"`
	TutorialId           string `json:"tutorialId"`
	Type                 string `json:"type"`
}

//20220531 leehs Redis 데이터를 가져와 저장
func (this Quest) Init(w http.ResponseWriter, req *http.Request) []Quest {
	tmp := redis.GetValue(cli, "quest")
	res := tmp.Val()
	quest := []Quest{}
	json.Unmarshal([]byte(res), &quest)

	return quest
}

//20220531 leehs id로 값 찾기
func (this Quest) GetDataByID(w http.ResponseWriter, req *http.Request) *Quest {
	quest := this.Init(w, req)
	id := req.FormValue("id")

	for _, i := range quest {
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
func (this Quest) GetDataByName(w http.ResponseWriter, req *http.Request) []Quest {
	quest := this.Init(w, req)
	desc := req.FormValue("name")
	var tmp []Quest

	for _, i := range quest {
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
