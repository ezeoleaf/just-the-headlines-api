package models

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Newspaper struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
}

type Newspapers struct {
	Newspapers []Newspaper `json:"items"`
}

func GetNewspapers(db *sql.DB) Newspapers {
	sql := "SELECT n.id, n.name, c.name FROM newspaper n INNER JOIN country c ON(n.country_id = c.id)"
	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	result := Newspapers{}
	for rows.Next() {
		n := Newspaper{}
		e := rows.Scan(&n.ID, &n.Name, &n.Country)
		if e != nil {
			panic(e)
		}
		result.Newspapers = append(result.Newspapers, n)
	}

	return result
}
