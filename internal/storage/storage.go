package storage

import types "github.com/parthpati1102/Golang-pRest-API-Project/internal/type"

type Storage interface {
	CreateStudent(name string, email string, age int) (int64, error)
	GetStudentById(id int64) (types.Student, error)
	GetStudents() ([]types.Student, error)
	UpdateStudent(id int64, name, email string, age int) error
	DeleteStudent(id int64) error
}
