package models

import (
	"database/sql"
)

type Section struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	RSS       string `json:"rss"`
	Failed    bool   `json:"failed"`
	Newspaper string `json:"newspaper"`
}

type Sections struct {
	Sections []Section `json:"sections"`
}

func sectionsFromRows(rows *sql.Rows) Sections {
	sections := Sections{}
	for rows.Next() {
		s := Section{}
		e := rows.Scan(&s.ID, &s.Name, &s.RSS, &s.Failed, &s.Newspaper)

		if e != nil {
			panic(e)
		}

		sections.Sections = append(sections.Sections, s)
	}

	return sections
}

func GetSectionsByName(db *sql.DB, name string) Sections {
	sql := `SELECT s.id, s.name, s.rss, s.failed, n.name FROM section s
	INNER JOIN newspaper n ON s.newspaper_id = n.id
	WHERE UPPER(s.name) LIKE '%' || $1 || '%'`

	rows, err := db.Query(sql, name)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	return sectionsFromRows(rows)
}

func GetSectionsByNewspaper(db *sql.DB, newspaper_id int) Sections {
	sql := `SELECT id, name, rss, failed, "" FROM section WHERE newspaper_id=$1`

	rows, err := db.Query(sql, newspaper_id)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	return sectionsFromRows(rows)
}

func GetSectionsByTag(db *sql.DB, tag string) {

}
