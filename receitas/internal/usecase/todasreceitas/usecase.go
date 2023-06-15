package todasreceitas

import (
	"receitas/internal/entity"
	"receitas/internal/repository"
)

type todasReceitas struct {
	repository repository.Interactor
}

func (t *todasReceitas) BuscarReceitas() []entity.Receita {
	receitas, _ := t.repository.FindAll()
	return receitas
}
