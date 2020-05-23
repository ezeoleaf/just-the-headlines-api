package models

import (
	"database/sql"
	"fmt"
)

// Section contains information related to the section table in the database
type Section struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	RSS        string `json:"rss"`
	Failed     bool   `json:"failed"`
	Newspaper  string `json:"newspaper"`
	Subscribed bool   `json:"subscribed"`
}

// Sections is a list of Section
type Sections struct {
	Sections []Section `json:"sections"`
}

func sectionsFromRows(rows *sql.Rows) Sections {
	sections := Sections{}
	for rows.Next() {
		s := Section{}
		e := rows.Scan(&s.ID, &s.Name, &s.RSS, &s.Failed, &s.Newspaper, &s.Subscribed)

		if e != nil {
			panic(e)
		}

		sections.Sections = append(sections.Sections, s)
	}

	return sections
}

func getSectionsByUser(db *sql.DB, userID int64) Sections {
	rows, err := db.Query(getUserSections, userID)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	return sectionsFromRows(rows)
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
func GetSectionsByNewspaper(db *sql.DB, newspaperID int, userID int64) Sections {
	rows, err := db.Query(SectionsByNewspaper, newspaperID, userID)
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

func Subscribe(db *sql.DB, sectionID int64, userID int64, subscribe bool) (int64, error) {
	var sql string
	if subscribe {
		sql = subscribeUser
	} else {
		sql = unsubscribeUser
	}

	s, e := db.Prepare(sql)
	if e != nil {
		panic(e)
	}

	defer s.Close()

	r, e := s.Exec(sectionID, userID)

	if e != nil {
		panic(e)
	}

	return r.RowsAffected()
}
