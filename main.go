package main

import (
	"fmt"
	"net/http"
	"time"

	"hwrap/internal/command"

	"github.com/labstack/echo/v4"
)

type (
	Response struct {
		Message string `json:"message"`
		Date    string `json:"date"`
	}
)

func main() {
	e := echo.New()
	Route(e)
	e.Logger.Fatal(e.Start(":8080"))
}

func Route(e *echo.Echo) {
	e.GET("/healthcheck", func(c echo.Context) error {
		return c.JSON(http.StatusOK, Response{
			Message: "success",
			Date:    time.Now().Format("2006-01-02T00:00:00"),
		})
	})

	e.GET("/ls", func(c echo.Context) error {
		args := c.QueryParam("args")
		f(args)

		fmt.Println(command.Exec())

		return c.JSON(http.StatusOK, Response{
			Message: "ls2",
			Date:    time.Now().Format("2006-01-02T00:00:00"),
		})
	})
}

func f(args string) {
	fmt.Println(args)
}
