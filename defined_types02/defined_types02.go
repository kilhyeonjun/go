/*
OOP : JAVA, C++, C# 등...
클래스안에 선언된 함수를 메서드라고 한다.

Go : 리시버(receiver) parameter
*/

package main

import "fmt"

type Number int

// func (n Number) Square() Number {
func (n *Number) Square() { // method
	*n = *n * *n
	// return n
}

func pointerSquare(n *int) { // function
	*n = *n * *n
}

func main() {
	n1 := Number(5)
	n2 := 4
	// fmt.Println(n1.Square())

	pointerSquare(&n2)
	fmt.Println(n2)

	fmt.Println(n1)
	n1.Square()
	fmt.Println(n1)
}

// package main

// import "fmt"

// type MyType string

// func (m MyType) myReturn() int {
// 	return len(m)
// }

// func main() {
// 	t1 := MyType("메서드테스트")
// 	t2 := MyType("단 축")
// 	fmt.Println(t1.myReturn())
// 	fmt.Println(t2.myReturn())
// }

// package main

// import "fmt"

// type MyType string

// func (m MyType) sayHi(n int, f bool) {
// 	fmt.Println(n, f)
// 	fmt.Println("Test : " + m)
// }

// func main() {
// 	t1 := MyType("메서드 테스트 중")
// 	t2 := MyType("단축 연산자도 사용가능")
// 	t1.sayHi(0, true)
// 	t2.sayHi(-7, false)
// }

// package main

// import "fmt"

// type (
// 	Liters      float64
// 	Gallons     float64
// 	Milliliters float64
// )

// func GallonsToLiters(g Gallons) Liters {
// 	return Liters(g * 3.785)
// }

// func GallonsToMilliliters(g Gallons) Liters {
// 	return Liters(g * 3785.41)
// }

// func LitersToGallons(l Liters) Gallons {
// 	return Gallons(l * 0.264)
// }

// func MillilitersToGallons(ml Milliliters) Gallons {
// 	return Gallons(ml * 0.000264)
// }

// // func ToLiters(g Gallons) Liters {
// // 	return Liters(g * 3.785)
// // }

// // func ToGallons(l Liters) Gallons {
// // 	return Gallons(l * 0.264)
// // }

// // func ToGallons(ml Milliliters) Gallons {
// // 	return Gallons(ml * 0.000264)
// // }

// func main() {
// 	var carFuel Liters = 10
// 	var busFuel Gallons = 60

// 	carFuel = carFuel + GallonsToLiters(Gallons(10.0))
// 	busFuel = busFuel + LitersToGallons(Liters(10.0))

// 	fmt.Println(carFuel, busFuel)
// }
