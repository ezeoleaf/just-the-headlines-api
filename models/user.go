package models

import (
	"database/sql"
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"

	"golang.org/x/crypto/bcrypt"
)

type user struct {
	ID       int
	Username string
	Password string
}

func hashPassword(p string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(p), 14)

	return string(bytes), err
}

func validatePasswordHash(p, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(p))

	return err == nil
}

func LoginUser(db *sql.DB, username string, password string) (string, error) {
	u := user{}

	row := db.QueryRow(loginUser, username)
	err := row.Scan(&u.ID, &u.Username, &u.Password)

	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}

	logged := validatePasswordHash(password, u.Password)

	if !logged {
		return "", errors.New("Bad credentials")
	}

	token := generateToken(u)

	return token, nil
}

func generateToken(u user) string {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = u.ID
	claims["username"] = u.Username
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(os.Getenv("SECRET_WORD")))

	if err != nil {
		panic(err)
	}

	return t
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
