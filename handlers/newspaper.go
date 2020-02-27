package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/ezeoleaf/just-the-headlines-api/models"
	"github.com/labstack/echo"
)

func GetNewspapers(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.GetNewspapers(db))
	}
}

func GetNewspaper(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		return c.JSON(http.StatusOK, models.GetNewspaper(db, id))
	}
}
