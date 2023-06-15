package api

import (
	"encoding/json"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"receitas/internal/entity"
	"receitas/internal/usecase/buscarreceita"
	"receitas/internal/usecase/novareceita"
	"receitas/internal/usecase/removereceita"
	"receitas/internal/usecase/todasreceitas"
	"strconv"
)

type handle struct {
	novareceita   novareceita.UseCase
	removereceita removereceita.UseCase
	todasreceitas todasreceitas.UseCase
	buscarreceita buscarreceita.UseCase
}

func (h *handle) NovaReceita(c echo.Context) error {
	reader := c.Request().Body
	decoder := json.NewDecoder(reader)
	var receita entity.Receita
	err := decoder.Decode(&receita)
	if err != nil {
		return err
	}
	receita, err = h.novareceita.Criar(receita)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, receita)
}

func (h *handle) RemoverReceita(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return errors.New("id invalido")
	}
	err = h.removereceita.Remover(int64(id))
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}

func (h *handle) TodasReceitas(c echo.Context) error {
	return c.JSON(
		http.StatusOK,
		h.todasreceitas.BuscarReceitas(),
	)
}

func (h *handle) BuscarReceita(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return errors.New("id invalido")
	}
	receita, err := h.buscarreceita.Buscar(int64(id))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, receita)
}
