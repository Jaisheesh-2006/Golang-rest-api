//* This is the concrete implementation of the Storage interface using SQLite as the database backend.

package sqlite

import (
	"database/sql"

	"github.com/Jaisheesh-2006/Golang-rest-api/internal/config"
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
