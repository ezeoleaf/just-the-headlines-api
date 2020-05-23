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

func GetNewspapersByName(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		name := strings.ToUpper(c.Param("name"))
		return c.JSON(http.StatusOK, models.GetNewspapersByName(db, name))
	}
}

func GetNewspaper(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		userID, _ := getUserFromJWT(c)
		return c.JSON(http.StatusOK, models.GetNewspaper(db, id, userID))
	}
}
