package server

import (
	"echo/redis"
	"encoding/json"
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

func (this Monster) Init() []Monster {
	tmp := redis.GetValue(cli, "monster")
	res := tmp.Val()
	json.Unmarshal([]byte(res), &monster)

	return monster
}

type Qitem struct {
	Id      string `json:"id"`
	TextKor string `json:"textKor"`
	TextEng string `json:"textEng"`
	TextChn string `json:"textChn"`
}

func (this Qitem) Init() []Qitem {
	tmp := redis.GetValue(cli, "qitem")
	res := tmp.Val()
	json.Unmarshal([]byte(res), &qitem)

	return qitem
}

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

func (this Prop) Init() []Prop {
	tmp := redis.GetValue(cli, "prop")
	res := tmp.Val()
	json.Unmarshal([]byte(res), &prop)

	return prop
}

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

func (this Quest) Init() []Quest {
	tmp := redis.GetValue(cli, "quest")
	res := tmp.Val()
	json.Unmarshal([]byte(res), &quest)

	return quest
}

type StringQuest struct {
	Id          string `json:"id"`
	Description string `json:"description"`
}

func (this StringQuest) Init() []StringQuest {
	tmp := redis.GetValue(cli, "stringquest")
	res := tmp.Val()
	json.Unmarshal([]byte(res), &sq)

	return sq
}

type StringItem struct {
	Id      string `json:"id"`
	TextKor string `json:"textKor"`
	TextEng string `json:"textEng"`
	TextChn string `json:"textChn"`
}

func (this StringItem) Init() []StringItem {
	tmp := redis.GetValue(cli, "stringitem")
	res := tmp.Val()
	json.Unmarshal([]byte(res), &si)

	return si
}

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

func (this Dialogue) Init() []Dialogue {
	tmp := redis.GetValue(cli, "dialogue")
	res := tmp.Val()
	json.Unmarshal([]byte(res), &dia)

	return dia
}

type NPC struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	ShapeInfo string `json:"shapeInfo"`
	ShopType  string `json:"shopType"`
	Type      string `json:"type"`
}

func (this NPC) Init() []NPC {
	tmp := redis.GetValue(cli, "npc")
	res := tmp.Val()
	json.Unmarshal([]byte(res), &npc)

	return npc
}

type StringName struct {
	Id      string `json:"id"`
	TextKor string `json:"textKor"`
	TextEng string `json:"textEng"`
	TextChn string `json:"textChn"`
}

func (this StringName) Init() []StringName {
	tmp := redis.GetValue(cli, "stringname")
	res := tmp.Val()
	json.Unmarshal([]byte(res), &sn)

	return sn
}
