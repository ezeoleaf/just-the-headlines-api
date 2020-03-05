package models

import (
	"database/sql"
	"fmt"
)

func GetIDsFromTags(db *sql.DB, tags string) []int {

	rows, err := db.Query(TagsIDs, tags)

	if err != nil {
		panic(err)
	}

	fmt.Println(rows)

	return []int{}
}
