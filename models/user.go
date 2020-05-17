package models

import (
	"database/sql"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"

	_ "github.com/joho/godotenv/autoload"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	ID       int64
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
		return "", echo.NewHTTPError(http.StatusNonAuthoritativeInfo, "Please provide valid credentials") //errors.New("Bad credentials")
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

func AttachFilter(db *sql.DB, userID int64, filterID int64) (int64, error) {
	var userFilterID int64
	var filterErr error
	row := db.QueryRow(getUserFilter, userID, filterID)
	err := row.Scan(&userFilterID)

	if err == sql.ErrNoRows {
		smtm, e := db.Prepare(attachFilter)

		if e != nil {
			panic(e)
		}

		defer smtm.Close()

		r, e := smtm.Exec(userID, filterID)

		if e != nil {
			panic(e)
		}

		userFilterID, filterErr = r.LastInsertId()
	} else if err != nil {
		panic(err)
	}

	return userFilterID, filterErr
}

func DetachFilter(db *sql.DB, userID int64, filterID int64) (int64, error) {
	smtm, err := db.Prepare(detachFilter)

	if err != nil {
		panic(err)
	}

	defer smtm.Close()

	r, err := smtm.Exec(userID, filterID)

	if err != nil {
		panic(err)
	}

	return r.RowsAffected()
}
