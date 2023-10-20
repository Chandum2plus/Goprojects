package main

import "fmt"

type Students struct {
	name    string
	id      int
	college string
	course  string
}
type Teacher struct {
	tName   string
	faculty string
	id      int
}

// Access the Students (structure)
func (s *Students) student(name, college, course string, id int) {
	s.name = name
	s.college = college
	s.course = course
	s.id = id
}

// Display the student
func (s *Students) disp() {
	fmt.Println("Name  :-", s.name)
	fmt.Println("College :-", s.college)
	fmt.Println("Course :-", s.course)
	fmt.Println("Identity :-", s.id)
}

// Access the Teacher (structure)
func (t *Teacher) teacher(name, fac string, id int) {
	t.tName = name
	t.faculty = fac
	t.id = id
}

// Display the Teacher
func (t *Teacher) disp() {
	fmt.Println("Name :-", t.tName)
	fmt.Println("Faculty :-", t.faculty)
	fmt.Println("Identity :-", t.id)
}

type Disp interface {
	disp()
}

func main() {
	var s Students // instance of StudentStructure
	var t Teacher  // instance of TeacherStructure
	s.student("Chandu", "SRM", "MCA", 12)

	t.teacher("Subodh Sir", "Full Stack Development", 34)

	var d Disp // instance of DispInterface
	//fmt.Println("===== CALLING THE FUNCTION ======")
	fmt.Println("=== STUDENT ===")
	d = &s   // d is store the address of the StudentStructure
	d.disp() // show student interface

	fmt.Println("=== TEACHER ===")
	d = &t   // d is store the address of TeacherStructure
	d.disp() // show teacher interface

}
