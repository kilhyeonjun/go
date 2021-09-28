package mymath

import "fmt"

func MyAbs(n int) int {
	if n < 0 {
		n = n * -1
	}
	return n
}

func MyPower(base int, exponent int) int {
	result := 1
	for i := 1; i <= exponent; i++ {
		result = result * base
	}
	return result
}

func Test() {
	fmt.Println("내가 만든 수학 패키지!")
}