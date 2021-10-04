// 201744004 A반 길현준 2학년 C반 수업
package main

import (
	"fmt"
	"log"
	"os"
)

func isPrime(n int) (bool, error) { // 소수 판별 함수
	if n < 2 {
		return false, fmt.Errorf("%d는(은) 소수가 아닙니다~", n)
	}

	for i := 2; i < n; i++ {
		if n%i == 0 { // 약수가 존재하면
			return false, nil // return
		}
	}
	return true, nil // true 리턴이면 소수, false 소수 X
}

func primeRange(a int, b int) { // 구간별 소수 판별 함수
	var numbers []int // 소수를 담을 slice
	var sum int = 0   // 합계
	if a > b {        // a, b swap
		temp := a
		a = b
		b = temp
	}
	for i := a; i <= b; i++ { // 반복하여 isPrime 호출(소수 판별)
		p, err := isPrime(i) // 소수판별 함수 호출
		if err != nil {      // 에러가 있으면
			fmt.Println(err)
			os.Exit(0) // 시스템 종료
		}
		if p { // 소수이면
			numbers = append(numbers, i) // slice에 값 추가
			sum = sum + i                // 합계 계산
		}
	}
	fmt.Println(numbers)                                    // slice 출력
	fmt.Println("평균 :", float64(sum)/float64(len(numbers))) // 평균 출력
}

func main() {
	var n1, n2 int
	fmt.Print("정수 2개 입력 : ")
	_, err := fmt.Scanln(&n1, &n2) // 2개의 값 입력
	if err != nil {
		log.Fatal(err)
	}
	primeRange(n1, n2) // 소수 판별 호출
}
