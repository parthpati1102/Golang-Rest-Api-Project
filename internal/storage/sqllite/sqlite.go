package sqlite

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/parthpati1102/Golang-pRest-API-Project/internal/config"
)

type Sqlite struct {
	Db *sql.DB
}

func New(cfg *config.Config) (*Sqlite, error) {
	db, err := sql.Open("sqlite3", cfg.StoragePath)

	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS student (
	   id INTEGER PRIMARY KEY AUTOINCREMENT,
	   NAME TEXT,
	   EMAIL TEXT,
	   AGE INTEGER
	)`)

	if err != nil {
		return nil, err
	}

	return &Sqlite{
		Db: db,
	}, nil
}

func (s *Sqlite) CreateStudent(name string, email string, age int) (int64, error) {

	stmt, err := s.Db.Prepare("INSERT INTO STUDENT(name, email , age) VALUES (?, ?, ?)")

	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	//Orer is must important

	result, err := stmt.Exec(name, email, age)

	if err != nil {
		return 0, nil
	}

	lastId, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return lastId, nil
}
