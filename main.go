package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"
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

var (
	optPort = flag.Int("p", 8080, "port number")
)

func main() {
	ParseArgs()
	e := echo.New()
	Route(e)
	e.Logger.Fatal(e.Start(":" + strconv.Itoa(*optPort)))
}

func ParseArgs() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: hwrap [flags] command\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	if len(flag.Args()) > 0 {
		cmd = flag.Args()[0]
	} else {
		cmd = "echo"
	}
	fmt.Printf("command: %s, port: %d\n", cmd, *optPort)
}

func Route(e *echo.Echo) {
	e.GET("/healthcheck", func(c echo.Context) error {
		return c.JSON(http.StatusOK, Response{
			Message: "success",
			Date:    time.Now().Format("2006-01-02T00:00:00"),
		})
	})

	e.GET("/", func(c echo.Context) error {
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
