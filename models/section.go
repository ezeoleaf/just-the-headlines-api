package models

import (
	"database/sql"
)

type Section struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	RSS    string `json:"rss"`
	Failed bool   `json:"failed"`
}

type Sections struct {
	Sections []Section `json:"sections"`
}

func GetSections(db *sql.DB, id int) Sections {
	sql := "SELECT id, name, rss, failed FROM section WHERE newspaper_id=$1"

	rows, err := db.Query(sql, id)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	sections := Sections{}
	for rows.Next() {
		s := Section{}
		e := rows.Scan(&s.ID, &s.Name, &s.RSS, &s.Failed)

		if e != nil {
			panic(e)
		}

		sections.Sections = append(sections.Sections, s)
	}

	return sections
}
