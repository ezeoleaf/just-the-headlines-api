package models

import (
	"database/sql"
	"fmt"
)

// Section contains information related to the section table in the database
type Section struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	RSS       string `json:"rss"`
	Failed    bool   `json:"failed"`
	Newspaper string `json:"newspaper"`
}

// Sections is a list of Section
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

// GetSectionsByName returns an instance of Sections filtered by section name
func GetSectionsByName(db *sql.DB, name string) Sections {
	rows, err := db.Query(SectionsByName, name)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	return sectionsFromRows(rows)
}

// GetSectionsByNewspaper returns an instance of Sections filtered by newspaperID
func GetSectionsByNewspaper(db *sql.DB, newspaperID int) Sections {
	rows, err := db.Query(SectionsByNewspaper, newspaperID)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	return sectionsFromRows(rows)
}

// GetSectionsByTags returns an instances of Sections that are mapped to a tag
func GetSectionsByTags(db *sql.DB, tags string) Sections {
	tagIDs := GetIDsFromTags(db, tags)

	fmt.Println(tagIDs)

	return Sections{}
}
