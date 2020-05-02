package models

import (
	"database/sql"
	"strings"
)

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
