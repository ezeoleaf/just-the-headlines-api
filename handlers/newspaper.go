package handlers

import (
	"database/sql"
	"net/http"
	"strconv"
	"strings"

	"github.com/ezeoleaf/just-the-headlines-api/models"
	"github.com/labstack/echo"
)

func GetNewspapers(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.GetNewspapers(db))
	}
}

func GetNewspapersByCountry(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		country_code := strings.ToUpper(c.Param("code"))
		return c.JSON(http.StatusOK, models.GetNewspapersByCountry(db, country_code))
	}
}

func GetNewspaper(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		return c.JSON(http.StatusOK, models.GetNewspaper(db, id))
	}
}
