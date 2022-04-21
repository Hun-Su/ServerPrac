//package main
//
//import (
//	"io/ioutil"
//	"net/http"
//	"strconv"
//	"strings"
//	"unicode"
//)
//
//func main() {
//	http.Handle("/", new(testHandler))
//
//	http.ListenAndServe("127.0.0.1:5000", nil)
//}
//
//type testHandler struct {
//	http.Handler
//}
//
//func (h *testHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
//	b, _ := ioutil.ReadAll(req.Body)
//	defer req.Body.Close()
//
//	var operator string
//	var a []string
//	var res string
//
//	for _, val := range string(b) {
//		if !unicode.IsDigit(val) {
//			operator = string(val)
//		}
//	}
//
//	a = strings.Split(string(b), operator)
//
//	v1, _ := strconv.Atoi(a[0])
//	v2, _ := strconv.Atoi(a[1])
//
//	switch operator {
//	case "+":
//		res = strconv.Itoa(v1 + v2)
//	case "-":
//		res = strconv.Itoa(v1 - v2)
//	case "*":
//		res = strconv.Itoa(v1 * v2)
//	case "/":
//		res = strconv.Itoa(v1 / v2)
//	default:
//		res = "Not a binary equation"
//	}
//
//	w.Write([]byte(res))
//}
