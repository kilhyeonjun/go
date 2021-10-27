package main

import "fmt"

type Student struct {
	Name  string
	Id    int
	Score float64
	Class string
	Level int
}

type ScholarshipStudent struct {
	// StudentInfo Student
	Student
	Level       int
	Scholarship int
}

func main() {
	s1 := ScholarshipStudent{
		Student{"홍길동", 12345678, 3.91, "A", 4},
		1,
		1000000,
	}
	fmt.Printf("장학금 수여 학생은 %s이고 장학금액은 %d원 입니다\n", s1.Name, s1.Scholarship)
	fmt.Println(s1.Student.Level)
}

// package main

// import "fmt"

// type Student struct {
// 	Name  string
// 	Id    int
// 	Score float64
// 	Class string
// }

// type ScholarshipStudent struct {
// 	// Name        string
// 	// Id          int
// 	// Score       float64
// 	// Class       string
// 	StudentInfo Student
// 	Level       int
// 	Scholarship int
// }

// func main() {
// 	var s1 ScholarshipStudent = ScholarshipStudent{
// 		Student{"홍길동", 12345678, 3.91, "A"},
// 		1,
// 		1000000,
// 	}

// 	fmt.Println(s1)
// 	fmt.Println(s1.Scholarship)
// 	fmt.Println(s1.StudentInfo.Id)

// 	// var s1 Student = Student{"홍길동", 12345678, 3.91, "A"}
// 	// var s2 Student = Student{Name: "최길동", Id: 12341234, Class: "B", Score: 3.76}
// 	// var s3 Student = Student{"박길동", 11112222, 4.01, "C"}

// 	// printStudent(s[2])
// 	// printStudent(s[0])
// 	// printStudent(s[1])
// }
