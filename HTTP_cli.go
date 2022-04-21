//package main
//
//import (
//	"fmt"
//	"strconv"
//	"strings"
//	"unicode"
//)
//
//func main() {
//	var s string
//	fmt.Scanln(&s)
//	var operator string
//	var a []string
//	var res string
//
//	for _, val := range s {
//		if !unicode.IsDigit(val) {
//			operator = string(val)
//		}
//	}
//
//	a = strings.Split(s, operator)
//
//	fmt.Println(a)
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
//	fmt.Println(res)
//}
