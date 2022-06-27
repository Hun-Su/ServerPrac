package HTTP

import (
	"echo/logging"
	"echo/server"
	"net/http"
	"reflect"
	"strings"
)

type TestHandler struct {
	resource    server.Resource
	dialogue    server.Dialogue
	npc         server.NPC
	qitem       server.Qitem
	stringquest server.StringQuest
	stringitem  server.StringItem
	stringname  server.StringName
	quest       server.Quest
	monster     server.Monster
	prop        server.Prop
}

//leehs 20220516 path와 메소드를 매핑하여 저장하는 map
var functions map[string]interface{}

//leehs 20220516 핸들러의 모든 필드들의 메소드들을 각각 map에 저장하는 함수
func (h *TestHandler) Init() {
	elem := reflect.ValueOf(*h)

	for i := 0; i < elem.NumField(); i++ { //
		f := elem.Field(i)
		ty := f.Type()
		cn := strings.ToLower(ty.Name())

		v := reflect.New(ty)

		for i := 0; i < v.NumMethod(); i++ {
			t := v.Type().Method(i)
			m := v.MethodByName(t.Name)
			name := t.Name

			if strings.ToUpper(name[:1]) != name[:1] {
				continue
			}

			mn := strings.ToLower(name)

			functions["/"+cn+"/"+mn] = m.Interface()
		}
	}
}

func initFunctions() {
	functions = make(map[string]interface{})

	list := TestHandler{}
	rp := server.ResourceProvider{}
	list.Init()
	rp.Init()
}

func (h TestHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	//initFunctions()
	//path := req.URL.Path
	//
	////leehs 20220516 functions에 저장된 데이터를 기준으로 path에 맞는 함수 호출
	//if f, i := functions[path]; i {
	//	_, err := Call(f, w, req)
	//	if err != nil {
	//		log.Println(err)
	//	}
	//} else {
	//	w.WriteHeader(404)
	//	log.Println("No such method")
	//}
	buff := make([]byte, 8192)
	str, err := req.Body.Read(buff)
	if err != nil {
		logging.LogInfo(err.Error())
	}
	w.Write(buff[:str])
}

//leehs 20220520 함수 호출
func Call(function interface{}, arg ...interface{}) (result []reflect.Value, err error) {
	f := reflect.ValueOf(function)

	in := make([]reflect.Value, len(arg))
	for i, param := range arg {
		in[i] = reflect.ValueOf(param)
	}
	result = f.Call(in)
	return
}
