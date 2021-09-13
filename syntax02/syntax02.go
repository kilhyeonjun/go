package main

import (
	"fmt"
	//	"reflect"
	"strings"
)

func main() {
	texts := "G@ M@ney~~"
	fmt.Println(texts)
	r := strings.NewReplacer("@","o")
	newText := r.Replace(texts)
	fmt.Println(newText)

	// // 변수명은 영문자로 시작해야된다.
	// // 영문 대문자의 경우 다른 패키지에서 접근할 수 있다. 소문자로 시작하는 변수는 동일 패키지에서만 접근 가능
	// // 관례로는
	// // camelCase 방식
	// // index 는 i 로 줄여서 선언
	// // zero value
	// var e string
	// var d bool
	// var c rune
	// var b float64
	// var a int
	// // a := 7

	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)
	// fmt.Println(d)
	// fmt.Println(e)
	
	// fmt.Printf("%T\n", c)
	// fmt.Printf("%T\n", a)

	// fmt.Println(reflect.TypeOf(c))
	// fmt.Println(reflect.TypeOf(a))
}