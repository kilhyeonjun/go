package main

import "fmt"

type Student struct {
	Name  string
	Id    int
	Score float64
	Class string
}

func printStudent(s Student) {
	fmt.Println("성명 : ", s.Name)
	fmt.Println("분반 : ", s.Class)
	fmt.Println("학번 : ", s.Id)
	fmt.Printf("평점 :  %.2f\n", s.Score)
}

// func defaultStudent(id int) Student {
// 	var s Student
// 	s.Id = id
// 	s.Name = "noname"
// 	s.Class = "001"
// 	s.Score = 0.0
// 	return s
// }

func defaultStudent(s *Student) {
	s.Id = 11112222
	s.Name = "noname"
	s.Class = "001"
	s.Score = 0.0
}

func main() {
	var s1 Student = Student{"홍길동", 12345678, 3.91, "A"}
	var s2 Student = Student{Id: 12341234, Class: "B"}
	// s3 := defaultStudent(11112222)
	// s3.Class = "C"
	var s3 Student
	defaultStudent(&s3)
	printStudent(s3)

	printStudent(s1)
	printStudent(s2)
}
