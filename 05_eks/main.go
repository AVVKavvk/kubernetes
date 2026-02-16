package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	fmt.Println("Hello world from image")

	count := 0

	e := echo.New()

	e.GET("/health", func(c echo.Context) error {

		count++
		return c.String(http.StatusOK, fmt.Sprintf("%d", count))
	})

	e.GET("/users", func(c echo.Context) error {

		count++
		return c.JSON(http.StatusOK, map[string]map[string]interface{}{"users": {
			"1": "user1",
			"2": "user2",
			"3": "user3",
		}})
	})

	e.Use(middleware.RequestLogger())

	e.Logger.Fatal(e.Start(":8000"))

}
