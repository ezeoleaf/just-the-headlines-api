package handlers

import (
	"database/sql"
	"net/http"

	"github.com/ezeoleaf/just-the-headlines-api/models"
	"github.com/labstack/echo"
)

func PostFilter(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Get user id from JWT
		filter := c.FormValue("filter")

		userID, _ := getUserFromJWT(c)

		id, err := models.PostFilter(db, filter, userID)

		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, H{
			"created": id,
		})
	}
}
