package main

import (
	"fmt"
)

//pass by pointer(call by pointer)
func swap(n1 *int, n2 *int){
	temp :=*n1
	*n1 = *n2
	*n2 = temp
}

func main(){
	// 포인터 : 메모리 주소를 값으로 가지는 타입. & 주소연산자, * 포인터기호
	var a int = 10
	b := 20


	c := &a
	fmt.Printf("%T\n",c)

	fmt.Println(a, b)
	swap(&a, &b)
	fmt.Println(a, b)
}