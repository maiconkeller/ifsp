package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func hello(c echo.Context) error {
	//return c.String(http.StatusOK, "hello ECHO!")
	return c.JSON(http.StatusOK, Response{
		Message: "Hello ECHO!",
	})
}

func main() {
	e := echo.New()
	e.GET("hello", hello)
	e.Logger.Fatal(e.Start(":8082"))
}
