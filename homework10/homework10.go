package main

import "fmt"

type coin map[string]int // 코인 거래소 key 코인명 value 코인 값
// map은 reference 타입

func (c coin) change(s string, i int) { // 코인 가격 변동 메소드
	c[s] = i
	fmt.Printf("%s이(가) 변동되었습니다.\n", s)
	c.list()
}

func (c coin) append(m map[string]int) { // 코인 추가 메소드
	for k, v := range m {
		c[k] = v
		fmt.Printf("%s이(가) 상장되었습니다.\n", k)
	}
	c.list()
}

// func (c coin) append(s string, i int) {
// 	c[s] = i
// 	fmt.Printf("%s이(가) 상장되었습니다.\n", s)
// }

func (c coin) list() { // 코인 목록 메소드
	for k, v := range c {
		fmt.Printf("%s는 %d원\n", k, v)
	}
	fmt.Printf("\n\n")
}

func main() {
	upBit := make(coin)                          // upBit 거래소 생성
	prevList := map[string]int{"비트코인": 73327000} // 기존 코인 목록
	upBit.append(prevList)                       // 기존 코인 목록 추가
	// upBit.append("도지", 322)
	newList := map[string]int{"도지": 322, "비트토렌트": 4} // 상장 예정 코인
	upBit.append(newList)                            // 코인 상장
	upBit.change("도지", 400)                          // 코인 가격 변동
}
