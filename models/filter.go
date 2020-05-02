package models

import (
	"database/sql"
	"strings"
)

type Filter struct {
	ID     int64  `json:"id"`
	Filter string `json:"name"`
}

type FilterCollection struct {
	Filters []Filter `json:"filters"`
}

func GetFilters(db *sql.DB) FilterCollection {
	rows, err := db.Query(getFilters)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	result := FilterCollection{}
	for rows.Next() {
		f := Filter{}
		e := rows.Scan(&f.ID, &f.Filter)

		if e != nil {
			panic(e)
		}

		result.Filters = append(result.Filters, f)
	}

	return result
}

func PostFilter(db *sql.DB, filter string, userID int64) (int64, error) {
	var filterID int64

	row := db.QueryRow(searchFilter, filter)
	err := row.Scan(&filterID)

	if err == sql.ErrNoRows {
		smtm, e := db.Prepare(createFilter)

		if e != nil {
			panic(e)
		}

		defer smtm.Close()

		r, e := smtm.Exec(strings.ToUpper(filter))

		if e != nil {
			panic(e)
		}

		filterID, _ = r.LastInsertId()
	} else if err != nil {
		panic(err)
	}

	_, err = AttachFilter(db, userID, filterID)

	if err != nil {
		panic(err)
	}

	return filterID, nil
}
