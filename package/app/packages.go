package main

import (
	"fmt"
	"gowork/packages/mymath"
)

func main(){
	mymath.Test()
	fmt.Println(mymath.MyPower(2,10))
	fmt.Println(mymath.MyPower(3,3))
	fmt.Println("HI")
	fmt.Println(mymath.MyAbs(-99))
}