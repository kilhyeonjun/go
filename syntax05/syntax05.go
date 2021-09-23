package main

import (
	"fmt"
	"log"
)

// 입력(ftm 패키지의 ScanLn())된 수의 소수 판정 프로그램 v0.8
func main() {
	var number int
	fmt.Print("정수 입력 : ")
	_, err := fmt.Scanln(&number)

	//fmt.Println(n)
	if err != nil{
		log.Fatal(err)
	}

	isPrime := true
	for i := 2; i < number; i++ { 
		if number % i == 0{
			isPrime = false
			break
		}
	}
	if isPrime {
		fmt.Println(number ,"는(은) 소수입니다!")
	}else {
		fmt.Println(number ,"는(은) 소수가 아닙니다~")
	}
}


// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// // 입력된 수의 소수 판정 프로그램 v0.7
// func main() {
// 	fmt.Print("정수 입력 : ")
// 	rd := bufio.NewReader(os.Stdin)
// 	in, err := rd.ReadString('\n')
	
// 	if err != nil { // 에러가 발생하면
// 		log.Fatal(err)
// 	}	

// 	in = strings.TrimSpace(in)
// 	number, err := strconv.Atoi(in)
// 	if err != nil { // 에러가 발생하면
// 		log.Fatal(err)
// 	}

// 	isPrime := true
// 	for i := 2; i < number; i++ { 
// 		if number % i == 0{
// 			isPrime = false
// 			break
// 		}
// 	}
// 	if isPrime {
// 		fmt.Println(number ,"는(은) 소수입니다!")
// 	}else {
// 		fmt.Println(number ,"는(은) 소수가 아닙니다~")
// 	}
// }