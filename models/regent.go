package model

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func GetDB() (*sql.DB, error) {
	var err error
	db, err = sql.Open("mysql", "root:Kepandean@95@tcp(localhost:3306)/geography")
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

type Regent struct {
	ID      int    `json:"idregent"`
	Name    string `json:"name"`
	Capital string `json:"capital"`
	Area    string `json:"area"`
}

func (r Regent) GetRegents(db *sql.DB) ([]Regent, error) {
	rows, err := db.Query("SELECT idregent, name, capital, area FROM regent")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var regents []Regent
	for rows.Next() {
		var r Regent
		err := rows.Scan(&r.ID, &r.Name, &r.Area, &r.Capital)
		if err != nil {
			return nil, err
		}
		regents = append(regents, r)
	}
	return regents, nil
}

func (r *Regent) GetRegent(db *sql.DB, id int) (*Regent, error) {
	row := db.QueryRow("SELECT * FROM regent WHERE idregent = ?", id)
	if row == nil {
		return nil, sql.ErrNoRows
	}

	err := row.Scan(&r.ID, &r.Name, &r.Area, &r.Capital)
	if err != nil {
		return nil, err
	}
	return r, nil
}
