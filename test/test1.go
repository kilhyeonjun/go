package main

import "fmt"

func c1(number int) []int {
	s1 := []int{}
	for i := 2; i < number; i++ {
		if number%i == 0 {
			// i
			s1 = append(s1, i)
		}
	}
	return s1
}

func main() {
	s2 := c1(100)
	fmt.Println(s2)
}
