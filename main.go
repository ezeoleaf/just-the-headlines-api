package main

import (
	"database/sql"
	"net/http"
	"os"

	"github.com/ezeoleaf/just-the-headlines-api/handlers"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	_ "github.com/mattn/go-sqlite3"
)

func startServer() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	storageName := os.Getenv("DB_NAME")

	db := initDB(storageName)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	initRoutes(e, db)

	e.Logger.Fatal(e.Start(":1323"))
}

func initDB(fp string) *sql.DB {
	db, err := sql.Open(os.Getenv("DB_DRIVER"), fp)

	if err != nil {
		panic(err)
	}

	if db == nil {
		panic("No DB created")
	}

	migrate(db)

	return db
}

func initRoutes(e *echo.Echo, db *sql.DB) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "JTH API")
	})
	e.GET("/newspapers", handlers.GetNewspapers(db))
	e.GET("/newspapers/:id", handlers.GetNewspaper(db))
	e.GET("/newspapers/country/:code", handlers.GetNewspapersByCountry(db))
	e.GET("/newspapers/name/:name", handlers.GetNewspapersByName(db))

	e.GET("/news/:id", handlers.GetNews(db))
	e.GET("/news/:id/:filter", handlers.GetFilteredNews(db))
	e.GET("/news/multiple/:sections/:filter", handlers.GetFilteredMultipleNews(db))
	e.GET("/news/multiple/:sections", handlers.GetMultipleNews(db))

	e.GET("/sections/name/:name", handlers.GetSectionsByName(db))
}

func migrate(db *sql.DB) {
	sql := `
	CREATE TABLE IF NOT EXISTS country(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name VARCHAR NOT NULL,
		code VARCHAR NOT NULL
	);
	CREATE TABLE IF NOT EXISTS newspaper(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name VARCHAR NOT NULL,
		country_id INTEGER REFERENCES country(id)
	);
	CREATE TABLE IF NOT EXISTS section(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name VARCHAR NOT NULL,
		rss VARCHAR NOT NULL,
		failed BOOLEAN DEFAULT FALSE,
		newspaper_id INTEGER REFERENCES newspaper(id)	
	);
	CREATE TABLE IF NOT EXISTS category(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name VARCHAR NOT NULL UNIQUE
	);
	CREATE TABLE IF NOT EXISTS tag(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		section_id INTEGER NOT NULL REFERENCES section(id),
		category_id INTEGER NOT NULL REFERENCES category(id)
	);
	`

	_, err := db.Exec(sql)

	if err != nil {
		panic(err)
	}
}

func main() {
	startServer()
}
