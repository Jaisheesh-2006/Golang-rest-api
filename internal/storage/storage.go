package storage

type Storage interface{
	CreateStudent(name string, age int, email string) (int64,error)
}