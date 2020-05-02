package main

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/middleware"
)

var IsLoggedIn = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte(os.Getenv("SECRET_WORD")),
})
