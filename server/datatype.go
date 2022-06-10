package server

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

type Qitem struct {
	Id      string `json:"id"`
	TextKor string `json:"textKor"`
	TextEng string `json:"textEng"`
	TextChn string `json:"textChn"`
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

type StringQuest struct {
	Id          string `json:"id"`
	Description string `json:"description"`
}

type StringItem struct {
	Id      string `json:"id"`
	TextKor string `json:"textKor"`
	TextEng string `json:"textEng"`
	TextChn string `json:"textChn"`
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

type NPC struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	ShapeInfo string `json:"shapeInfo"`
	ShopType  string `json:"shopType"`
	Type      string `json:"type"`
}

type StringName struct {
	Id      string `json:"id"`
	TextKor string `json:"textKor"`
	TextEng string `json:"textEng"`
	TextChn string `json:"textChn"`
}
