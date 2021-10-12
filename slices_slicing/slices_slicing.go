package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3}
	// s2 := make([]int, len(s1))

	// for i, v := range s1 {
	// 	s2[i] = v
	// }

	// s2 := append([]int{}, s1[0], s1[1], s1[2])

	s2 := append([]int{}, s1...)

	fmt.Println(s1, s2)
	fmt.Printf("%x %x\n", &s1[0], &s2[0])
}

// package main

// import "fmt"

// func main() {
// 	// a := []string{"a", "b", "c", "d"}
// 	// as := a[0:2]
// 	// as[1] = "z"

// 	// b := [4]int{4, 3, 2, 1}
// 	// bs := b[1:3]
// 	// fmt.Println(bs)

// 	a := make([]string, 4, 5)
// 	// a := []string{"a", "b", "c", "d"}
// 	a[0] = "a"
// 	a[2] = "c"
// 	a[3] = "d"
// 	as := a[0:2]
// 	as[1] = "z"
// 	// c := append(a, "y")
// 	c := append(a, "x", "y")

// 	c = append(c, "1", "2", "3", "4", "5")

// 	fmt.Println(a, len(a), cap(a))
// 	fmt.Println(c, len(c), cap(c))
// 	fmt.Printf("%x %x %x", &a[0], &as[0], &c[0])
// }
