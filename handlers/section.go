package handlers

import (
	"database/sql"
	"net/http"
	"strings"

	"github.com/ezeoleaf/just-the-headlines-api/models"
	"github.com/labstack/echo"
)

func GetSectionsByName(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		name := strings.ToUpper(c.Param("name"))
		return c.JSON(http.StatusOK, models.GetSectionsByName(db, name))
	}
}
