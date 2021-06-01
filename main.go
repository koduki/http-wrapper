package main

import (
	"fmt"
	"net/http"
	"strings"
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

var cmd string

func main() {
	cmd = "echo"

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
		hargs := c.QueryParam("args")
		args := strings.Split(hargs, ",")

		fmt.Printf("cmd: %s, args: %s\n", cmd, args)
		status := command.Exec(cmd, args)
		fmt.Printf("status: %s\n", status)

		return c.JSON(http.StatusOK, Response{
			Message: status,
			Date:    time.Now().Format("2006-01-02T00:00:00"),
		})
	})
}
