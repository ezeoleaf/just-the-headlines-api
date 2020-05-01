package models

import (
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

func hashPassword(p string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(p), 14)

	return string(bytes), err
}

func validatePasswordHash(p, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(p))

	return err == nil
}

func PostUser(db *sql.DB, email string, username string, password string) (int64, error) {
	smtm, err := db.Prepare(newUser)

	if err != nil {
		panic(err)
	}

	defer smtm.Close()

	p, err := hashPassword(password)

	if err != nil {
		panic(err)
	}

	r, err := smtm.Exec(email, username, p)

	if err != nil {
		panic(err)
	}

	return r.LastInsertId()
}
