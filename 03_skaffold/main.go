package main

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "6000"
	}
	e := echo.New()

	e.Use(middleware.RequestLogger())

	e.GET("/", func(ctx echo.Context) error {

		return ctx.JSON(200, "Hello World")
	})

	e.GET("/new/test", func(ctx echo.Context) error {

		return ctx.JSON(200, "New Hello World")
	})

	e.Logger.Fatal(e.Start(":" + port))

}
