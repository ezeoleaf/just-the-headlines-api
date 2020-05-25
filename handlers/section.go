package handlers

import (
	"database/sql"
	"net/http"
	"strconv"
	"strings"

	"github.com/ezeoleaf/just-the-headlines-api/models"
	"github.com/labstack/echo"
)

func GetSectionsByName(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		name := strings.ToUpper(c.Param("name"))
		userID, _ := getUserFromJWT(c)
		return c.JSON(http.StatusOK, models.GetSectionsByName(db, name, userID))
	}
}

func Subscribe(db *sql.DB, subscribe bool) echo.HandlerFunc {
	return func(c echo.Context) error {
		sectionID, _ := strconv.ParseInt(c.FormValue("sectionID"), 10, 64)
		userID, _ := getUserFromJWT(c)

		_, err := models.Subscribe(db, sectionID, userID, subscribe)

		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, H{
			"success": true,
		})
	}
}
