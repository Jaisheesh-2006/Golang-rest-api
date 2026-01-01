package storage

import "github.com/Jaisheesh-2006/Golang-rest-api/internal/types"

type Storage interface {
	CreateStudent(name string, age int, email string) (int64, error)
	GetStudentById(id int64) (types.Student, error)
	GetStudents() ([]types.Student, error)
	UpdateStudentById(id int64, name string, age int, email string) error
}