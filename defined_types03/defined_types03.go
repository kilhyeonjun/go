package main

import (
	"fmt"
	"time"
)

type student struct {
	id       int
	name     string
	subject  string
	seceonds int
}

func (s student) study() {
	fmt.Printf("%s는 %s과목을 공부합니다.", s.name, s.subject)
}

func (s student) countDown() {
	for s.seceonds > 0 {
		fmt.Println(s.seceonds)
		// time.Sleep(time.Second * 2)
		time.Sleep(time.Second)
		s.seceonds--
	}
}

func funcStudy(fs student) {
	fmt.Printf("%s는 %s과목을 공부합니다.", fs.name, fs.subject)
}

func main() {
	s1 := student{1000, "홍길동", "Go", 5}
	// fmt.Println(s1)
	s1.countDown()
	s1.study() // method
	// funcStudy(s1) // function
}

// package main

// import "fmt"

// type (
// 	Liters      float64
// 	Gallons     float64
// 	Milliliters float64
// )

// func (g Gallons) ToLiters() Liters {
// 	return Liters(g * 3.785)
// }

// func (g Gallons) ToMilliliters() Liters {
// 	return Liters(g * 3785.41)
// }

// func (l Liters) ToGallons() Gallons {
// 	return Gallons(l * 0.264)
// }

// func (ml Milliliters) ToGallons() Gallons {
// 	return Gallons(ml * 0.000264)
// }

// func main() {
// 	// coke := Liters(2)
// 	var coke Liters = 2
// 	fanta := Milliliters(750)

// 	fmt.Printf("%.2f리터는 %0.3f갤론과 같습니다\n", coke, coke.ToGallons())
// 	fmt.Printf("%.2f밀리리터는 %0.3f갤론과 같습니다\n", fanta, fanta.ToGallons())
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

// func main() {
// 	var carFuel Liters = 10
// 	var busFuel Gallons = 60

// 	carFuel = carFuel + GallonsToLiters(Gallons(10.0))
// 	busFuel = busFuel + LitersToGallons(Liters(10.0))

// 	fmt.Println(carFuel, busFuel)
// }
