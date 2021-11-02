package main

import "fmt"

type (
	Liters  float64
	Gallons float64
)

func ToLiters(g Gallons) Liters {
	return Liters(g * 3.785)
}

func ToGallons(l Liters) Gallons {
	return Gallons(l * 0.264)
}

func main() {
	var carFuel Liters = 10
	var busFuel Gallons = 60

	carFuel = carFuel + ToLiters(Gallons(10.0))
	busFuel = busFuel + ToGallons(Liters(10.0))

	fmt.Println(carFuel, busFuel)
}

// package main

// import "fmt"

// type (
// 	Liters  float64
// 	Gallons float64
// )
// type Name string

// func main() {
// 	// var carfuel Liters = 10
// 	// var busfuel Gallons = 60
// 	carfuel := Liters(10)
// 	busfuel := Gallons(60)

// 	carfuel = Liters(Gallons(10) * 3.785)
// 	busfuel = Gallons(Liters(60) * 0.264)

// 	// carfuel = Liters(15.9)
// 	// carfuel = Gallons(15.9)
// 	// busfuel = Gallons(50.5)
// 	fmt.Println(carfuel, busfuel)
// 	fmt.Println("defined types")

// 	fmt.Println(Name("홍길동") + Name("홍길순"))
// 	// fmt.Println(Name("홍길동") - Name("홍길순"))
// 	fmt.Println(Name("홍길동") > Name("홍길순"))
// 	fmt.Println(Name("홍길동") < Name("홍길순"))

// 	// fmt.Println(Gallons(1.2) + Gallons(5.5))
// 	// fmt.Println(Gallons(1.2) - Gallons(5.5))
// 	// fmt.Println(Gallons(1.2) / Gallons(5.5))
// 	// fmt.Println(Gallons(1.2) * Gallons(5.5))
// 	// fmt.Println(Gallons(1.2) == Gallons(5.5))
// 	// fmt.Println(Gallons(1.2) != Gallons(5.5))
// 	// fmt.Println(Gallons(1.2) > Gallons(5.5))
// 	// fmt.Println(Gallons(1.2) < Gallons(5.5))
// }
