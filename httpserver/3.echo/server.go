package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "hello ECHO!")
}

func main() {
	e := echo.New()
	e.GET("hello", hello)
	e.Logger.Fatal(e.Start(":8081"))
}
