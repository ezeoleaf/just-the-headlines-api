package main

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/middleware"
)

var isLoggedIn = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte(os.Getenv("SECRET_WORD")),
})
