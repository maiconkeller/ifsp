package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"receitas/internal/controller/api"
)

func Start(port string, handle api.Controller) {
	// instancia do Echo
	server := echo.New()

	// middlewares
	server.Use(middleware.Logger())
	server.Use(middleware.Recover())
	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{AllowCredentials: true}))

	// URI prefixada
	group := server.Group("/v1")

	// rotas
	group.GET("healtz", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})
	group.POST("/receitas", handle.NovaReceita)
	group.GET("/receitas", handle.TodasReceitas)
	group.DELETE("/receitas/:id", handle.RemoverReceita)
	group.GET("/receitas/:id", handle.BuscarReceita)

	//start
	server.Logger.Fatal(server.Start(port))

}
