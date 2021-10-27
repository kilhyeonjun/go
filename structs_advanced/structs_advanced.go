package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Name  string
	Id    int
	Score float64
	Class string
}

type Students []Student

func (s Students) Len() int           { return len(s) }
func (s Students) Less(i, j int) bool { return s[i].Score < s[j].Score }
func (s Students) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func main() {
	s := []Student{
		{"홍길동", 12345678, 3.91, "A"},
		{Name: "최길동", Id: 12341234, Class: "B", Score: 3.76},
		{"박길동", 11112222, 4.01, "C"},
	}

	sort.Sort(Students(s))
	fmt.Println(s)
	// var s1 Student = Student{"홍길동", 12345678, 3.91, "A"}
	// var s2 Student = Student{Name: "최길동", Id: 12341234, Class: "B", Score: 3.76}
	// var s3 Student = Student{"박길동", 11112222, 4.01, "C"}

	// printStudent(s[2])
	// printStudent(s[0])
	// printStudent(s[1])
}

// package main

// import "fmt"

// type Student struct {
// 	Name  string
// 	Id    int
// 	Score float64
// 	Class string
// }

// func printStudent(s Student) {
// 	fmt.Println("성명 : ", s.Name)
// 	fmt.Println("분반 : ", s.Class)
// 	fmt.Println("학번 : ", s.Id)
// 	fmt.Printf("평점 :  %.2f\n", s.Score)
// }

// // func defaultStudent(id int) Student {
// // 	var s Student
// // 	s.Id = id
// // 	s.Name = "noname"
// // 	s.Class = "001"
// // 	s.Score = 0.0
// // 	return s
// // }

// func defaultStudent(s *Student) {
// 	s.Id = 11112222
// 	s.Name = "noname"
// 	s.Class = "001"
// 	s.Score = 0.0
// }

// func main() {
// 	var s1 Student = Student{"홍길동", 12345678, 3.91, "A"}
// 	var s2 Student = Student{Id: 12341234, Class: "B"}
// 	// s3 := defaultStudent(11112222)
// 	// s3.Class = "C"
// 	var s3 Student
// 	defaultStudent(&s3)
// 	printStudent(s3)

// 	printStudent(s1)
// 	printStudent(s2)
// }
