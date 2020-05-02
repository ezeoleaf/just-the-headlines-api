package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/ezeoleaf/just-the-headlines-api/models"
	"github.com/labstack/echo"
)

// H is a type used to return data
type H map[string]interface{}

// PostUser saves an user to the database
func PostUser(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		username := c.FormValue("username")
		email := c.FormValue("email")
		password := c.FormValue("password")

		id, err := models.PostUser(db, email, username, password)

		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, H{
			"created": id,
		})
	}
}

func LoginUser(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		username := c.FormValue("username")
		password := c.FormValue("password")

		t, err := models.LoginUser(db, username, password)

		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, H{
			"token": t,
		})
	}
}

func UserFilter(db *sql.DB, attach bool) echo.HandlerFunc {
	return func(c echo.Context) error {
		filterID, _ := strconv.ParseInt(c.FormValue("filterID"), 10, 64)
		userID, _ := getUserFromJWT(c)

		var id int64
		var err error

		if attach {
			id, err = models.AttachFilter(db, userID, filterID)
		} else {
			id, err = models.DetachFilter(db, userID, filterID)
		}

		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, H{
			"success": true,
			"id":      id,
		})
	}
}
