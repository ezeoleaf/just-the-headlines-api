package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/ezeoleaf/just-the-headlines-api/models"
	"github.com/labstack/echo"
)

func GetNewsBySection(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		return c.JSON(http.StatusOK, models.GetNewsBySection(db, id))
	}
}

func GetFilteredNews(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		filter := c.Param("filter")
		return c.JSON(http.StatusOK, models.GetFilteredNews(db, id, filter))
	}
}

func GetMultipleNews(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		sections := c.Param("sections")
		return c.JSON(http.StatusOK, models.GetMultipleNews(db, sections, ``))
	}
}

func GetFilteredMultipleNews(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		sections := c.Param("sections")
		filter := c.Param("filter")
		return c.JSON(http.StatusOK, models.GetMultipleNews(db, sections, filter))
	}
}
