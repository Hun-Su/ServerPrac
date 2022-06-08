package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//20220531 leehs id로 값 찾기
func (this Dialogue) GetDataByID(w http.ResponseWriter, req *http.Request) *Dialogue {
	id := req.FormValue("id")

	for _, i := range dia {
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
	Type := req.FormValue("type")
	var tmp []Dialogue

	for _, i := range dia {
		if i.Type == Type {
			tmp = append(tmp, i)
		}
	}
	for _, i := range tmp {
		tmp1, _ := json.MarshalIndent(i, "", " ")
		fmt.Println(string(tmp1))
	}
	if len(tmp) == 0 {
		log.Println("Type not found")
	}
	return tmp
}

//20220531 leehs id로 값 찾기
func (this Monster) GetDataByID(w http.ResponseWriter, req *http.Request) *Monster {
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
	name := req.FormValue("name")
	var tmp []Monster

	for _, i := range monster {
		if i.Name == name {
			tmp = append(tmp, i)
		}
	}
	for _, i := range tmp {
		tmp1, _ := json.MarshalIndent(i, "", " ")
		fmt.Println(string(tmp1))
	}
	if len(tmp) == 0 {
		log.Println("Name not found")
	}
	return tmp
}

//20220531 leehs id로 값 찾기
func (this NPC) GetDataByID(w http.ResponseWriter, req *http.Request) *NPC {
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
	Type := req.FormValue("type")
	var tmp []NPC

	for _, i := range npc {
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
		log.Println("Type not found")
	}
	return tmp
}

//20220531 leehs id로 값 찾기
func (this Prop) GetDataByID(w http.ResponseWriter, req *http.Request) *Prop {
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
	name := req.FormValue("name")
	var tmp []Prop

	for _, i := range prop {
		if i.Name == name {
			tmp = append(tmp, i)
		}
	}
	for _, i := range tmp {
		tmp1, _ := json.MarshalIndent(i, "", " ")
		fmt.Println(string(tmp1))
	}
	if len(tmp) == 0 {
		log.Println("Name not found")
	}
	return tmp
}

//20220531 leehs id로 값 찾기
func (this Qitem) GetDataByID(w http.ResponseWriter, req *http.Request) *Qitem {
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

//20220531 leehs id로 값 찾기
func (this Quest) GetDataByID(w http.ResponseWriter, req *http.Request) *Quest {
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
	name := req.FormValue("name")
	var tmp []Quest

	for _, i := range quest {
		if i.Name == name {
			tmp = append(tmp, i)
		}
	}
	for _, i := range tmp {
		tmp1, _ := json.MarshalIndent(i, "", " ")
		fmt.Println(string(tmp1))
	}
	if len(tmp) == 0 {
		log.Println("Name not found")
	}
	return tmp
}

//20220531 leehs id로 값 찾기
func (this StringItem) GetDataByID(w http.ResponseWriter, req *http.Request) *StringItem {
	id := req.FormValue("id")

	for _, i := range si {
		if i.Id == id {
			tmp, _ := json.MarshalIndent(i, "", " ")
			fmt.Println(string(tmp))
			return &i
		}
	}
	log.Println("ID not found")
	return nil
}

//20220531 leehs id로 값 찾기
func (this StringName) GetDataByID(w http.ResponseWriter, req *http.Request) *StringName {
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

//20220531 leehs id로 값 찾기
func (this StringQuest) GetDataByID(w http.ResponseWriter, req *http.Request) *StringQuest {
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
		log.Println("Desc not found")
	}
	return tmp
}
