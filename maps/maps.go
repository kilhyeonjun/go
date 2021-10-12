package main

import "fmt"

func status(name string) {
	balls := map[string]int{"성기훈": 20, "오일남": 0}
	ball, exists := balls[name]
	if !exists {
		// fmt.Println(name, "남은", balls[name], "개로 탈락")
		fmt.Println(name, "남은 게임 참여자가 아닙니다")
	} else if ball < 1 {
		fmt.Println(name, "남은", balls[name], "개로 탈락")
	} else {
		fmt.Println(name, "남은 게임에서 승리하셨습니다.")
	}
}

func main() {
	status("오일남")
	status("강철")
	status("성기훈")
}

// package main

// import "fmt"

// func main() {
// 	// balls := make(map[string]int)

// 	var balls map[string]int
// 	// balls = make(map[string]int)

// 	fmt.Println(balls)

// 	balls["성기훈"] = 20
// 	fmt.Println(balls["성기훈"])
// 	fmt.Println(balls["오일남"])
// }

// package main

// import "fmt"

// func main() {
// 	// map literal
// 	games := map[int]string{456: "성기훈", 218: "박해수", 67: "강새벽", 1: "오일남", 199: "알리", 101: "아이오아이"}

// 	// fmt.Println(games[067])

// 	// 입력순서와 상관없이 랜덤
// 	for _, v := range games {
// 		fmt.Println(v)
// 	}

// 	games[101] = "장덕수" // update
// 	delete(games, 199)

// 	for k, v := range games {
// 		fmt.Println(k, v)
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	// var games map[int]string
// 	// games = make(map[int]string)

// 	games := make(map[int]string)

// 	// append
// 	games[456] = "성기훈"
// 	games[218] = "박해수"
// 	games[067] = "강새벽"
// 	games[001] = "오일남"
// 	games[199] = "알리"
// 	games[101] = "아이오아이"

// 	// fmt.Println(games[067])

// 	// 입력순서와 상관없이 랜덤
// 	for _, v := range games {
// 		fmt.Println(v)
// 	}

// 	games[101] = "장덕수" // update
// 	delete(games, 199)

// 	for k, v := range games {
// 		fmt.Println(k, v)
// 	}
// }
