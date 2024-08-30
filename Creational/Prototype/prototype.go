package prototype

import "fmt"

// Интерфейс прототипа студента
type StudentPrototyper interface {
	Clone() StudentPrototyper
	GetFullName() string
	GetAge() int
	ChangeFullName(string, string)
}

// Конкретный студент
type Student struct {
	Name    string
	Surname string
	Age     int
}

func NewStudent(name, surname string, age int) StudentPrototyper {
	return &Student{
		Name:    name,
		Surname: surname,
		Age:     age,
	}
}

func (s Student) GetFullName() string {
	fn := fmt.Sprintf("%s %s", s.Name, s.Surname)
	return fn
}

func (s Student) GetAge() int {
	return s.Age
}

func (s *Student) ChangeFullName(name, surname string) {
	s.Name = name
	s.Surname = surname
}

// Возвращает новый объект студента
func (s *Student) Clone() StudentPrototyper {
	return &Student{s.Name, s.Surname, s.Age}
}
