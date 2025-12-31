//* This is the concrete implementation of the Storage interface using SQLite as the database backend.

package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/Jaisheesh-2006/Golang-rest-api/internal/config"
	"github.com/Jaisheesh-2006/Golang-rest-api/internal/types"
	_ "github.com/mattn/go-sqlite3"
)

type SqLite struct {
	Db *sql.DB
}

// * constructor function to initialize SqLite struct
func New(cfg *config.Config) (*SqLite, error) {
	db, err := sql.Open("sqlite3", cfg.StoragePath)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
	   CREATE TABLE IF NOT EXISTS students (
	     id INTEGER PRIMARY KEY AUTOINCREMENT,
		 name TEXT,
		 email TEXT,  
		 age INTEGER
	   )
	`)
	if err != nil {
		return nil, err
	}
	return &SqLite{Db: db}, nil
}

func (s *SqLite) CreateStudent(name string, age int, email string) (int64, error) {
	stmt, err := s.Db.Prepare("INSERT INTO students (name, age, email) VALUES (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(name, age, email)
	if err != nil {
		return 0, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return lastId, nil

}

func (s *SqLite) GetStudentById(id int64) (types.Student, error) {
	stmt, err := s.Db.Prepare("SELECT * FROM students WHERE id = ? LIMIT 1")
	if err != nil {
		return types.Student{}, err
	}
	defer stmt.Close()

	var student types.Student

	err = stmt.QueryRow(id).Scan(&student.Id, &student.Name, &student.Email, &student.Age)

	if err != nil {
		if err == sql.ErrNoRows {
			return types.Student{}, fmt.Errorf("No student found with id %d %w", id, err)
		}
		return types.Student{}, fmt.Errorf("Query error %w", err)
	}

	return student, nil

}
func (s *SqLite) GetStudents() ([]types.Student, error)  {
	stmt, err := s.Db.Prepare("SELECT * FROM students")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []types.Student

	for rows.Next() {
		var student types.Student
		err := rows.Scan(&student.Id, &student.Name, &student.Email, &student.Age)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}
	return students, nil
	
}