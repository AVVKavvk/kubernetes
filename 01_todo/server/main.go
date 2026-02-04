package main

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "5001"
	}

	e := echo.New()

	e.Use(middleware.RequestLogger())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	todo := e.Group("/todos")

	todo.POST("", CreateTodo)

	todo.GET("", GetTodos)

	todo.PUT("/:id", MarkTodoDone)

	todo.DELETE("/:id", DeleteTodo)

	e.Logger.Fatal(e.Start(":" + port))

}
