package main

import (
	"github.com/everestmz/fuzzing-demo/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.POST("/rle", handlers.EncodeHandler)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
