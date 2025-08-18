package main

import "fmt"

//- **통합 과제:** 학생 정보 관리 프로그램 만들기
//1.  `Student`라는 이름의 `struct`를 정의하세요. 이 구조체는 `name` (string), `id` (int), `grades` ([3]int 타입의 배열) 필드를 가져야 합니다.
//2.  여러 `Student` 객체를 저장할 수 있는 `slice`를 생성하세요.
//3.  새로운 학생을 추가하는 `addStudent` 함수를 작성하세요. 이 함수는 학생 목록(slice)과 학생 정보를 인자로 받습니다.
//4.  특정 학생의 성적을 수정하는 `updateGrade` 함수를 작성하세요. 이 함수는 학생 객체에 대한 `pointer`를 인자로 받아 해당 학생의 성적을 변경해야 합니다.
//5.  모든 학생의 정보를 출력하는 `printAllStudents` 함수를 작성하세요.
//6.  `main` 함수에서 위 함수들을 호출하여 2명 이상의 학생을 추가하고, 그중 한 명의 성적을 수정한 뒤, 전체 학생 정보를 출력하는 코드를 완성하세요.

type Student struct {
	name   string
	id     int
	grades [3]int
}

func addStudent(studentList []Student, student Student) []Student {
	studentList = append(studentList, student)
	return studentList
}

func updateGrade(studentList []Student, id int, index int, grade int) []Student {
	for i := range studentList {
		if studentList[i].id == id {
			studentList[i].grades[index] = grade
		}
	}
	return studentList
}

func printAllStudents(studentList []Student) {
	for i := 0; i < len(studentList); i++ {
		fmt.Println(studentList[i])
	}
}

func main() {
	var studentList []Student
	studentList = addStudent(studentList, Student{"John", 1, [3]int{90, 70, 80}})
	studentList = addStudent(studentList, Student{"Bill", 2, [3]int{20, 30, 10}})
	printAllStudents(studentList)

	studentList = updateGrade(studentList, 2, 2, 40)
	printAllStudents(studentList)
}
