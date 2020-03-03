package models

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Newspaper struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Country     string    `json:"country"`
	CountryCode string    `json:"contry_code"`
	Sections    []Section `json:"sections"`
}

type Newspapers struct {
	Newspapers []Newspaper `json:"items"`
}

func GetNewspapers(db *sql.DB) Newspapers {
	sql := "SELECT n.id, n.name, c.name, c.code FROM newspaper n INNER JOIN country c ON(n.country_id = c.id)"
	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	return newspapersFromRows(rows)
}

func newspapersFromRows(rows *sql.Rows) Newspapers {
	result := Newspapers{}
	for rows.Next() {
		n := Newspaper{}
		e := rows.Scan(&n.ID, &n.Name, &n.Country, &n.CountryCode)
		if e != nil {
			panic(e)
		}
		result.Newspapers = append(result.Newspapers, n)
	}

	return result
}

func GetNewspapersByName(db *sql.DB, name string) Newspapers {
	sql := `SELECT n.id, n.name, c.name, c.code FROM newspaper n INNER JOIN country c ON(n.country_id = c.id) WHERE UPPER(n.name) LIKE '%' || $1 || '%'`

	rows, err := db.Query(sql, name)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	return newspapersFromRows(rows)
}

func GetNewspapersByCountry(db *sql.DB, country_code string) Newspapers {
	sql := `SELECT n.id, n.name, c.name, c.code FROM newspaper n INNER JOIN country c ON(n.country_id = c.id) WHERE UPPER(c.code) LIKE $1 || '%'`

	rows, err := db.Query(sql, country_code)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	return newspapersFromRows(rows)
}

func GetNewspaper(db *sql.DB, id int) Newspaper {
	sql := "SELECT n.id, n.name, c.name, c.code FROM newspaper n INNER JOIN country c ON(n.country_id = c.id) WHERE n.id=$1"
	n := Newspaper{}

	row := db.QueryRow(sql, id)
	e := row.Scan(&n.ID, &n.Name, &n.Country, &n.CountryCode)

	if e != nil {
		panic(e)
	}

	sections := GetSections(db, id)

	n.Sections = sections.Sections

	return n
}
