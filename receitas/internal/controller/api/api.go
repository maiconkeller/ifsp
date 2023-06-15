package api

import (
	"github.com/labstack/echo/v4"
	"receitas/internal/usecase/buscarreceita"
	"receitas/internal/usecase/novareceita"
	"receitas/internal/usecase/removereceita"
	"receitas/internal/usecase/todasreceitas"
)

type Controller interface {
	NovaReceita(c echo.Context) error
	RemoverReceita(c echo.Context) error
	TodasReceitas(c echo.Context) error
	BuscarReceita(c echo.Context) error
}

func New(novareceita novareceita.UseCase, removereceita removereceita.UseCase, todasreceitas todasreceitas.UseCase,
	buscarreceita buscarreceita.UseCase) Controller {
	return &handle{
		novareceita:   novareceita,
		removereceita: removereceita,
		todasreceitas: todasreceitas,
		buscarreceita: buscarreceita,
	}
}
