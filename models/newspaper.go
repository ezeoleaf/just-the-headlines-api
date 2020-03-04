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
	rows, err := db.Query(NewspapersAll)

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
	rows, err := db.Query(NewspapersByName, name)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	return newspapersFromRows(rows)
}

func GetNewspapersByCountry(db *sql.DB, countryCode string) Newspapers {
	rows, err := db.Query(NewspapersByCountry, countryCode)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	return newspapersFromRows(rows)
}

func GetNewspaper(db *sql.DB, id int) Newspaper {
	n := Newspaper{}

	row := db.QueryRow(NewspaperByID, id)
	e := row.Scan(&n.ID, &n.Name, &n.Country, &n.CountryCode)

	if e != nil {
		panic(e)
	}

	sections := GetSectionsByNewspaper(db, id)

	n.Sections = sections.Sections

	return n
}
