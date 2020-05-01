package handlers

import (
	"database/sql"
	"net/http"

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
