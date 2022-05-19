package HTTP

import (
	"echo/server"
	"log"
	"net/http"
	"reflect"
	"strings"
)

type TestHandler struct {
	resource server.Resource
}

//leehs 20220516 핸들러의 모든 필드들의 메소드들을 각각 map에 저장하는 함수
func (h *TestHandler) init() {

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

//leehs 20220516 path와 메소드를 매핑하여 저장하는 map
var functions map[string]interface{}

func initFunctions() {
	functions = make(map[string]interface{})

	list := TestHandler{}
	list.init()
}

func (h TestHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	initFunctions()

	path := req.URL.Path
	param := strings.Split(req.FormValue("name"), ",")
	var in []reflect.Value
	for _, i := range param {
		if i != "" {
			in = append(in, reflect.ValueOf(i))
		}
	}

	//leehs 20220516 functions에 저장된 데이터를 기준으로 path에 맞는 함수 호출
	if f, i := functions[path]; i {
		ff := reflect.ValueOf(f)
		if len(in) == 0 {
			ff.Call(nil)
		} else {
			ff.Call(in)
		}
	} else {
		w.WriteHeader(404)
		log.Println("No such method")
	}
	//}
	//b, _ := ioutil.ReadAll(req.Body)
	//defer req.Body.Close()
	//
	//var operator string
	//var a []string
	//var res string
	//
	//for _, val := range string(b) {
	//	if !unicode.IsDigit(val) {
	//		operator = string(val)
	//	}
	//}
	//
	//a = strings.Split(string(b), operator)
	//
	//v1, _ := strconv.Atoi(a[0])
	//v2, _ := strconv.Atoi(a[1])
	//
	//switch operator {
	//case "+":
	//	res = strconv.Itoa(v1 + v2)
	//case "-":
	//	res = strconv.Itoa(v1 - v2)
	//case "*":
	//	res = strconv.Itoa(v1 * v2)
	//case "/":
	//	res = strconv.Itoa(v1 / v2)
	//default:
	//	res = "Not a binary equation"
	//}
	//
	//w.Write([]byte(res))
}
